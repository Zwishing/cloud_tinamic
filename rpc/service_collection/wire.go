// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	"cloud_tinamic/rpc/service_collection/repo"
	"github.com/google/wire"
)

func InitServiceCollection() (*ServiceCollectionImpl, error) {
	wire.Build(
		repo.NewDB,
		repo.NewServiceCollectionRepoImpl,
		NewMapProcessorClient,
		NewGeoServiceClient,
		NewSourceServiceClient,
		NewServiceCollectionImpl,
	)
	return nil, nil
}
