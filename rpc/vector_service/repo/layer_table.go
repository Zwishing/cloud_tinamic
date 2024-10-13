package repo

import (
	"cloud_tinamic/kitex_gen/service/vector"
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"

	// Database
	"cloud_tinamic/rpc/vector_service/repo/cql"

	"github.com/jackc/pgtype"
)

// LayerTable provides metadata about the table layer
type LayerTable struct {
	ID             string
	Schema         string
	Table          string
	Description    string
	Properties     map[string]TableProperty
	GeometryType   string
	IDColumn       string
	GeometryColumn string
	Srid           int
}

// TableProperty provides metadata about a single property field,
// features in a table layer may have multiple such fields
type TableProperty struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	order       int
}

/********************************************************************************
 * Layer Interface
 */

// GetType disambiguates between function and table layers
func (lyr *LayerTable) GetType() LayerType {
	return LayerTypeTable
}

// GetID returns the complete ID (schema.name) by which to reference a given layer
func (lyr *LayerTable) GetID() string {
	return lyr.ID
}

// GetDescription returns the text description for a layer
// or an empty string if no description is set
func (lyr *LayerTable) GetDescription() string {
	return lyr.Description
}

// GetName returns just the name of a given layer
func (lyr *LayerTable) GetName() string {
	return lyr.Table
}

// GetSchema returns just the schema for a given layer
func (lyr *LayerTable) GetSchema() string {
	return lyr.Schema
}

// GetTileQuery GetTileRequest takes tile and request parameters as input and returns a TileRequest
// specifying the SQL to fetch appropriate data
func (lyr *LayerTable) GetTileQuery(tile Tile, params *vector.QueryParameters) TileQuery {
	ReviseQueryParameters(params)
	sql, _ := lyr.requestSQL(&tile, params)

	tr := TileQuery{
		LayerID: lyr.ID,
		Tile:    tile,
		SQL:     sql,
		Args:    nil,
	}
	return tr
}

/********************************************************************************/

type queryParameters struct {
	Limit      int
	Properties []string
	Resolution int
	Buffer     int
	Filter     string
	FilterCrs  int
}

// getRequestPropertiesParameter compares the properties in the request
// with the properties in the table layer, and returns a slice of
// just those that occur in both, or a slice of all table properties
// if there is not query parameter, or no matches
func (lyr *LayerTable) getQueryPropertiesParameter(q url.Values) []string {
	sAtts := make([]string, 0)
	haveProperties := false

	for k, v := range q {
		if strings.EqualFold(k, "properties") {
			sAtts = v
			haveProperties = true
			break
		}
	}

	lyrAtts := (*lyr).Properties
	queryAtts := make([]string, 0, len(lyrAtts))
	haveIDColumn := false

	if haveProperties {
		aAtts := strings.Split(sAtts[0], ",")
		for _, att := range aAtts {
			decAtt, err := url.QueryUnescape(att)
			if err == nil {
				decAtt = strings.Trim(decAtt, " ")
				att, ok := lyrAtts[decAtt]
				if ok {
					if att.Name == lyr.IDColumn {
						haveIDColumn = true
					}
					queryAtts = append(queryAtts, att.Name)
				}
			}
		}
	}
	// No request parameter or no matches, so we want to
	// return all the properties in the table layer
	if len(queryAtts) == 0 {
		for _, v := range lyrAtts {
			queryAtts = append(queryAtts, v.Name)
		}
	}
	if (!haveIDColumn) && lyr.IDColumn != "" {
		queryAtts = append(queryAtts, lyr.IDColumn)
	}
	return queryAtts
}

// ReviseQueryParameters getRequestParameters reads user-settables parameters
// from the request URL, or uses the system defaults
// if the parameters are not set
func ReviseQueryParameters(params *vector.QueryParameters) {
	if params.Limit < 0 {
		params.Limit = -1
	}
	if params.Resolution < 0 {
		params.Resolution = 4096
	}
	if params.Buffer < 0 {
		params.Buffer = 256
	}
	if params.FilterCrs < 0 {
		params.FilterCrs = 4326
	}
}

/********************************************************************************/

// GetBoundsExact returns the data coverage extent for a table layer
// in EPSG:4326, clipped to (+/-180, +/-90)
func (lyr *LayerTable) GetBoundsExact() (Bounds, error) {
	bounds := Bounds{}
	extentSQL := fmt.Sprintf(`
	WITH ext AS (
		SELECT
			coalesce(
				ST_Transform(ST_SetSRID(ST_Extent("%s"), %d), 4326),
				ST_MakeEnvelope(-180, -90, 180, 90, 4326)
			) AS geom
		FROM "%s"."%s"
	)
	SELECT
		ST_XMin(ext.geom) AS xmin,
		ST_YMin(ext.geom) AS ymin,
		ST_XMax(ext.geom) AS xmax,
		ST_YMax(ext.geom) AS ymax
	FROM ext
	`, lyr.GeometryColumn, lyr.Srid, lyr.Schema, lyr.Table)

	var (
		xmin pgtype.Float8
		xmax pgtype.Float8
		ymin pgtype.Float8
		ymax pgtype.Float8
	)
	err := db.QueryRow(context.Background(), extentSQL).Scan(&xmin, &ymin, &xmax, &ymax)
	if err != nil {
		return bounds, fmt.Errorf("%s", err)
	}

	bounds.SRID = 4326
	bounds.Xmin = xmin.Float
	bounds.Ymin = ymin.Float
	bounds.Xmax = xmax.Float
	bounds.Ymax = ymax.Float
	bounds.sanitize()
	return bounds, nil
}

// GetBounds returns the estimated extent for a table layer, transformed to EPSG:4326
func (lyr *LayerTable) GetBounds() (Bounds, error) {
	bounds := Bounds{}
	extentSQL := fmt.Sprintf(`
		WITH ext AS (
			SELECT ST_Transform(ST_SetSRID(ST_EstimatedExtent('%s', '%s', '%s'), %d), 4326) AS geom
		)
		SELECT
			ST_XMin(ext.geom) AS xmin,
			ST_YMin(ext.geom) AS ymin,
			ST_XMax(ext.geom) AS xmax,
			ST_YMax(ext.geom) AS ymax
		FROM ext
		`, lyr.Schema, lyr.Table, lyr.GeometryColumn, lyr.Srid)

	var (
		xmin pgtype.Float8
		xmax pgtype.Float8
		ymin pgtype.Float8
		ymax pgtype.Float8
	)
	err := db.QueryRow(context.Background(), extentSQL).Scan(&xmin, &ymin, &xmax, &ymax)
	if err != nil {
		return bounds, fmt.Errorf("%s", err)
	}

	// Failed to get estimate? Get the exact bounds.
	if xmin.Status == pgtype.Null {
		warning := fmt.Sprintf("Estimated extent query failed, run 'ANALYZE %s.%s'", lyr.Schema, lyr.Table)
		klog.Warn(warning)
		return lyr.GetBoundsExact()
	}

	bounds.SRID = 4326
	bounds.Xmin = xmin.Float
	bounds.Ymin = ymin.Float
	bounds.Xmax = xmax.Float
	bounds.Ymax = ymax.Float
	bounds.sanitize()
	return bounds, nil
}

func (lyr *LayerTable) requestSQL(tile *Tile, qp *vector.QueryParameters) (string, error) {

	type sqlParameters struct {
		TileSQL        string
		QuerySQL       string
		FilterSQL      string
		TileSrid       int
		Resolution     int32
		Buffer         int32
		Properties     string
		MvtParams      string
		Limit          string
		Schema         string
		Table          string
		GeometryColumn string
		Srid           int
	}

	// need both the exact tile boundary for clipping and an
	// expanded version for querying
	tileBounds := tile.Bounds
	queryBounds := tile.Bounds
	queryBounds.Expand(tile.width() * float64(qp.Buffer) / float64(qp.Resolution))
	tileSQL := tileBounds.SQL()
	tileQuerySQL := queryBounds.SQL()

	filterSQL, err := lyr.filterSQL(qp)
	if err != nil {
		return "", err
	}

	// SRID of the tile we are going to generate, which might be different
	// from the layer SRID in the database
	tileSrid := tile.Bounds.SRID

	// preserve case and special characters in column names
	// of SQL query by double quoting names
	attrNames := make([]string, 0, len(qp.Properties))
	for _, a := range qp.Properties {
		attrNames = append(attrNames, fmt.Sprintf("\"%s\"", a))
	}

	// only specify MVT format parameters we have configured
	mvtParams := make([]string, 0)
	mvtParams = append(mvtParams, fmt.Sprintf("'%s', %d", lyr.ID, qp.Resolution))
	if lyr.GeometryColumn != "" {
		mvtParams = append(mvtParams, fmt.Sprintf("'%s'", lyr.GeometryColumn))
	}
	// The idColumn parameter is PostGIS3+ only
	if lyr.IDColumn != "" {
		mvtParams = append(mvtParams, fmt.Sprintf("'%s'", lyr.IDColumn))
	}

	sp := sqlParameters{
		TileSQL:        tileSQL,
		QuerySQL:       tileQuerySQL,
		FilterSQL:      filterSQL,
		TileSrid:       tileSrid,
		Resolution:     qp.Resolution,
		Buffer:         qp.Buffer,
		Properties:     strings.Join(attrNames, ", "),
		MvtParams:      strings.Join(mvtParams, ", "),
		Schema:         lyr.Schema,
		Table:          lyr.Table,
		GeometryColumn: lyr.GeometryColumn,
		Srid:           lyr.Srid,
	}

	if qp.Limit > 0 {
		sp.Limit = fmt.Sprintf("LIMIT %d", qp.Limit)
	}

	// TODO: Remove ST_Force2D when fixes to line clipping are common
	// in GEOS. See https://trac.osgeo.org/postgis/ticket/4690
	tmplSQL := `
	SELECT ST_AsMVT(mvtgeom, {{ .MvtParams }}) FROM (
		SELECT ST_AsMVTGeom(
			ST_Transform(ST_Force2D(t."{{ .GeometryColumn }}"), {{ .TileSrid }}),
			bounds.geom_clip,
			{{ .Resolution }},
			{{ .Buffer }}
		  ) AS "{{ .GeometryColumn }}"
		  {{ if .Properties }}
		  , {{ .Properties }}
		  {{ end }}
		FROM "{{ .Schema }}"."{{ .Table }}" t, (
			SELECT {{ .TileSQL }}  AS geom_clip,
					{{ .QuerySQL }} AS geom_query
			) bounds
		WHERE ST_Intersects(t."{{ .GeometryColumn }}",
							ST_Transform(bounds.geom_query, {{ .Srid }}))
			{{ .FilterSQL }}
		{{ .Limit }}
	) mvtgeom
	`

	sql, err := renderSQLTemplate("tabletilesql", tmplSQL, sp)
	if err != nil {
		return "", err
	}
	return sql, err
}

func (lyr *LayerTable) filterSQL(qp *vector.QueryParameters) (string, error) {
	filter := qp.Filter

	// 直接创建一个新的 strings.Builder 实例
	sqlBuilder := new(strings.Builder)

	// 使用预编译的正则表达式来优化字符串处理
	var filterRegex = regexp.MustCompile(`\s+`)
	filter = filterRegex.ReplaceAllString(filter, " ")

	// 使用并发处理来加速SQL转换
	errChan := make(chan error, 1)
	sqlChan := make(chan string, 1)

	go func() {
		sql, err := cql.TranspileToSQL(filter, int(qp.FilterCrs), lyr.Srid)
		if err != nil {
			errChan <- err
			return
		}
		sqlChan <- sql
	}()

	select {
	case err := <-errChan:
		close(sqlChan) // 关闭sqlChan以避免泄漏
		return "", err
	case sql := <-sqlChan:
		if sql != "" {
			sqlBuilder.WriteString("AND ")
			sqlBuilder.WriteString(sql)
		}
		return sqlBuilder.String(), nil // 返回构建的 SQL 字符串
	case <-time.After(5 * time.Second):
		return "", fmt.Errorf("filterSQL operation timed out")
	}
}
