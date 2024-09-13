// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	"cloud_tinamic/rpc/auth/repo"
	"github.com/google/wire"
)

func InitAuthService() *AuthServiceImpl {
	wire.Build(
		repo.NewPgPool,
		repo.NewEnforcer,
		repo.NewAuthRepoImpl,
		NewAuthServiceImpl,
	)
	return nil
}
