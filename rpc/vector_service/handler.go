package main

import (
	vector "cloud_tinamic/kitex_gen/service/vector"
	. "cloud_tinamic/pkg/errors"
	"cloud_tinamic/rpc/vector_service/repo"
	"context"
	"time"
)

// VectorServiceImpl implements the last service interface defined in the IDL.
type VectorServiceImpl struct {
	VectorServiceRepo repo.VectorServiceRepo
}

func NewVectorServiceImpl(repo repo.VectorServiceRepo) *VectorServiceImpl {
	return &VectorServiceImpl{
		VectorServiceRepo: repo,
	}
}

// GetTile implements the VectorServiceImpl interface.
func (s *VectorServiceImpl) GetTile(ctx context.Context, serviceKey string,
	x, y int32, z int8, ext string, params *vector.QueryParameters) (resp []byte, err error) {

	lyr, err := s.VectorServiceRepo.GetLayer(serviceKey)
	if err != nil {
		return nil, Kerrorf(NotFoundCode, "Layer not found for serviceKey: %s", serviceKey)
	}

	tile, err := repo.MakeTile(x, y, z, ext)
	if err != nil {
		return nil, Kerrorf(InvalidParametersCode, "Failed to create tile: %s", err.Error())
	}

	// Use a shorter timeout for high concurrency
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	tileQuery := lyr.GetTileQuery(tile, params)

	// Use connection pooling for database connections
	mvt, err := s.VectorServiceRepo.GetTileFromDB(ctx, &tileQuery)
	if err != nil {
		return nil, Kerrorf(DatabaseErrorCode, "Failed to get tile from DB: %s", err.Error())
	}

	return mvt, nil
}
