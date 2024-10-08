// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"cloud_tinamic/rpc/user/repo"
)

// Injectors from wire.go:

func InitUserService() *UserServiceImpl {
	db := repo.NewDB()
	userRepo := repo.NewUserRepoImpl(db)
	userServiceImpl := NewUserServiceImpl(userRepo)
	return userServiceImpl
}
