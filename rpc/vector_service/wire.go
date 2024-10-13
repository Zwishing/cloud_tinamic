// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	"cloud_tinamic/rpc/vector_service/repo"
	"github.com/google/wire"
)

func InitVectorService() *VectorServiceImpl {
	wire.Build(
		repo.NewPgPool,
		repo.InitCache,
		repo.NewVectorServiceRepoImpl,
		NewVectorServiceImpl,
	)
	return nil
}
