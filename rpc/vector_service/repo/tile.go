package repo

import (
	"fmt"
	"math"
)

// Tile represents a single tile in a tile
// pyramid, usually referenced in a URL path
// of the form "Zoom/X/Y.Ext"
type Tile struct {
	Zoom   int8   `json:"zoom"`
	X      int32  `json:"x"`
	Y      int32  `json:"y"`
	Ext    string `json:"ext"`
	Bounds Bounds `json:"bounds"`
}

// makeTile uses the map populated by the mux.Router
// containing x, y and z keys, and extracts integers
// from them
func MakeTile(x, y int32, zoom int8, ext string) (Tile, error) {
	tile := Tile{Zoom: zoom, X: x, Y: y, Ext: ext}
	e := tile.CalculateBounds()
	return tile, e
}

func (tile *Tile) width() float64 {
	return math.Abs(tile.Bounds.Xmax - tile.Bounds.Xmin)
}

// IsValid tests that the tile contains
// only tile addresses that fit within the
// zoom level, and that the zoom level is
// not crazy large
func (tile *Tile) IsValid() bool {
	if tile.Zoom > 32 || tile.Zoom < 0 {
		return false
	}
	worldTileSize := int32(1) << uint32(tile.Zoom)
	if tile.X < 0 || tile.X >= worldTileSize ||
		tile.Y < 0 || tile.Y >= worldTileSize {
		return false
	}
	return true
}

// CalculateBounds calculates the cartesian bounds that
// correspond to this tile
func (tile *Tile) CalculateBounds() (e error) {
	serverBounds, e := getServerBounds()
	if e != nil {
		return e
	}

	worldWidthInTiles := float64(int(1) << uint(tile.Zoom))
	tileWidth := math.Abs(serverBounds.Xmax-serverBounds.Xmin) / worldWidthInTiles

	// Calculate geographic bounds from tile coordinates
	// XYZ tile coordinates are in "image space" so origin is
	// top-left, not bottom right
	xmin := serverBounds.Xmin + (tileWidth * float64(tile.X))
	xmax := serverBounds.Xmin + (tileWidth * float64(tile.X+1))
	ymin := serverBounds.Ymax - (tileWidth * float64(tile.Y+1))
	ymax := serverBounds.Ymax - (tileWidth * float64(tile.Y))
	tile.Bounds = Bounds{serverBounds.SRID, xmin, ymin, xmax, ymax}

	return nil
}

// String returns a path-like representation of the Tile
func (tile *Tile) String() string {
	return fmt.Sprintf("%d/%d/%d.%s", tile.Zoom, tile.X, tile.Y, tile.Ext)
}
