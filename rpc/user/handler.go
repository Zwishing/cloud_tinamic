package main

import (
	user "cloud_tinamic/kitex_gen/base/user"
	"cloud_tinamic/rpc/user/repo"
	"context"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	UserRepo repo.UserRepo
}

func NewUserServiceImpl(repo repo.UserRepo) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepo: repo,
	}
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	return
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, req *user.InfoRequest) (resp *user.InfoResponse, err error) {
	// TODO: Your code here...
	return
}
