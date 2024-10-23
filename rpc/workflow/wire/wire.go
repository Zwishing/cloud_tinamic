// wire.go
//go:build wireinject
// +build wireinject

package wire

import (
	"cloud_tinamic/rpc/workflow"
	"github.com/google/wire"
)

func InitVectorActivities() (*workflow.VectorActivities, error) {
	wire.Build(
		workflow.NewGeoServiceClient,
		workflow.NewMapProcessorClient,
		workflow.NewVectorActivities,
	)
	return nil, nil
}
