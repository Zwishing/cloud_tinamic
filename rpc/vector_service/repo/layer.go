package repo

import (
	"cloud_tinamic/kitex_gen/service/vector"
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
	GetTileQuery(tile Tile, params *vector.QueryParameters) TileQuery
}

// A TileQuery specifies what to fetch from the database for a single tile
type TileQuery struct {
	LayerID string
	Tile    Tile
	SQL     string
	Args    []interface{}
}

type LayerJSON struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Schema      string `json:"schema"`
	Description string `json:"description"`
	DetailURL   string `json:"detailurl"`
}
