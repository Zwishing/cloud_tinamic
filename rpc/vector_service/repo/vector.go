package repo

import (
	. "cloud_tinamic/pkg/errors"
	"cloud_tinamic/pkg/pg"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/patrickmn/go-cache"
)

type VectorServiceRepo interface {
	GetLayer(serviceKey string) (Layer, error)
	GetTileFromDB(ctx context.Context, tr *TileQuery) ([]byte, error)
}

type VectorServiceRepoImpl struct {
	db    *pg.PgPool
	cache *cache.Cache
}

func NewVectorServiceRepoImpl(db *pg.PgPool, c *cache.Cache) VectorServiceRepo {
	return &VectorServiceRepoImpl{
		db:    db,
		cache: c,
	}
}

func (v *VectorServiceRepoImpl) GetLayer(serviceKey string) (Layer, error) {
	if lyr, ok := v.cache.Get(serviceKey); ok {
		return lyr.(*LayerTable), nil
	}

	layer := LayerTable{
		Schema:         "vector",
		Table:          serviceKey,
		IDColumn:       "id",
		GeometryColumn: "geom",
	}

	err := v.db.QueryRow(
		context.Background(),
		"SELECT v.id, i.sird, v.geometry_category FROM service.info i JOIN service.vector v ON i.source_key = v.source_key WHERE i.service_key = $1",
		serviceKey,
	).Scan(&layer.ID, &layer.Srid, &layer.GeometryType)

	if err != nil {
		return nil, Kerrorf(QueryFailedCode, "Failed to get layer for service key %s: %v", serviceKey, err)
	}
	
	v.cache.SetDefault(serviceKey, &layer)
	return &layer, nil
}

// GetTileFromDB 使用原生的db查询提高性能
func (v *VectorServiceRepoImpl) GetTileFromDB(ctx context.Context, tr *TileQuery) ([]byte, error) {
	row := v.db.QueryRow(ctx, tr.SQL, tr.Args...)
	var mvtTile []byte
	err := row.Scan(&mvtTile)
	if err != nil {
		klog.Warn(err)
		// check for errors retrieving the rendered tile from the database
		// Timeout errors can occur if the context deadline is reached
		// or if the context is canceled during/before a database query.
		if pgconn.Timeout(err) {
			return nil, Kerrorf(TimeoutCode, "Timeout: deadline exceeded on %s/%s", tr.LayerID, tr.Tile.String())
		}
		return nil, Kerrorf(QueryFailedCode, "SQL error on %s/%s", tr.LayerID, tr.Tile.String())
	}
	return mvtTile, nil
}
