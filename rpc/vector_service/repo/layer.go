package repo

import (
	"cloud_tinamic/kitex_gen/service/vector"
	. "cloud_tinamic/pkg/errors"
)

// LayerType is the table/function type of a layer
type LayerType int

const (
	// LayerTypeTable is a table layer
	LayerTypeTable = 1
	// LayerTypeFunction is a function layer
	LayerTypeFunction = 2
)

func (lt LayerType) String() string {
	switch lt {
	case LayerTypeTable:
		return "table"
	case LayerTypeFunction:
		return "function"
	default:
		return "unknown"
	}
}

// A Layer is a LayerTable or a LayerFunction
// in either case it should be able to generate
// a TileRequest containing SQL to produce tiles
// given an input tile
type Layer interface {
	GetType() LayerType
	GetID() string
	GetDescription() string
	GetName() string
	GetSchema() string
	GetTileRequest(tile Tile, params *vector.QueryParameters) TileRequest
}

// A TileRequest specifies what to fetch from the database for a single tile
type TileRequest struct {
	LayerID string
	Tile    Tile
	SQL     string
	Args    []interface{}
}

func GetLayer(lyrID string) (Layer, error) {
	lyr, ok := c.Get(lyrID)
	if ok {
		return lyr.(Layer), nil
	}
	return lyr.(Layer), Kerrorf(NotFoundCode, "Unable to get layer '%s'", lyrID)
}

type layerJSON struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Schema      string `json:"schema"`
	Description string `json:"description"`
	DetailURL   string `json:"detailurl"`
}
