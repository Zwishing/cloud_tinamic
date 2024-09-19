package main

import (
	base "cloud_tinamic/kitex_gen/base"
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
	resp = &auth.AuthResponse{
		Base:  &base.BaseResp{},
		Allow: s.AuthRepo.Auth(req.Sub, req.Obj, req.Act),
	}

	if resp.Allow {
		resp.Base.Msg = "通过认证"
		resp.Base.Code = base.Code_SUCCESS
	} else {
		resp.Base.Msg = "无权限访问"
		resp.Base.Code = base.Code_UNAUTHORIZED
	}

	return resp, nil
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
