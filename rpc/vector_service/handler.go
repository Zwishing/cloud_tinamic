package main

import (
	vector "cloud_tinamic/kitex_gen/service/vector"
	"context"
)

// VectorServiceImpl implements the last service interface defined in the IDL.
type VectorServiceImpl struct{}

// Publish implements the VectorServiceImpl interface.
func (s *VectorServiceImpl) Publish(ctx context.Context, req string) (err error) {
	// TODO: Your code here...
	return
}

// GetTile implements the VectorServiceImpl interface.
func (s *VectorServiceImpl) GetTile(ctx context.Context, req *vector.GetTileRequest) (resp []byte, err error) {
	// TODO: Your code here...
	return
}
