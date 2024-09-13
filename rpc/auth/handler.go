package main

import (
	auth "cloud_tinamic/kitex_gen/base/auth"
	"cloud_tinamic/rpc/auth/repo"
	"context"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct {
	AuthRepo repo.AuthRepo
}

func NewAuthServiceImpl(repo repo.AuthRepo) *AuthServiceImpl {
	return &AuthServiceImpl{
		AuthRepo: repo,
	}
}

// Auth implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) Auth(ctx context.Context, req *auth.AuthResquest) (resp *auth.AuthResponse, err error) {
	// TODO: Your code here...
	return
}

// AddPolicy implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) AddPolicy(ctx context.Context, req *auth.AuthResquest) (resp *auth.EditResponse, err error) {
	// TODO: Your code here...
	return
}

// RemovePolicy implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) RemovePolicy(ctx context.Context, req *auth.AuthResquest) (resp *auth.EditResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdatePolicy implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) UpdatePolicy(ctx context.Context, req *auth.AuthResquest) (resp *auth.EditResponse, err error) {
	// TODO: Your code here...
	return
}
