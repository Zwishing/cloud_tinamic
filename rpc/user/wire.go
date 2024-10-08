// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	"cloud_tinamic/rpc/user/repo"
	"github.com/google/wire"
)

func InitUserService() *UserServiceImpl {
	wire.Build(
		repo.NewDB,
		repo.NewUserRepoImpl,
		NewUserServiceImpl,
	)
	return nil
}
