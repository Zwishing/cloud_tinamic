package main

import (
	base "cloud_tinamic/kitex_gen/base"
	user "cloud_tinamic/kitex_gen/base/user"
	"cloud_tinamic/rpc/user/pack"
	"cloud_tinamic/rpc/user/repo"
	"context"
	"fmt"
	uuid2 "github.com/google/uuid"
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
	resp = user.NewRegisterResponse()
	resp.SetBase(base.NewBaseResp())

	uuid, _ := uuid2.NewUUID()

	account := user.NewAccount()
	account.SetUserId(uuid.String())
	account.SetUsername(req.Username)
	account.SetUserCategory(req.UserCategory)

	usr := user.NewUser()
	usr.SetUserId(uuid.String())
	salt := RandomSalt()
	usr.SetPassword(CreateHashPassword(req.Password, salt))
	usr.SetSalt(salt)

	success,_ := s.UserRepo.AddUser(account, usr)

	if !success {
		resp.Base.Code = base.Code_FAIL
		resp.Base.Msg = "注册失败"
		err = fmt.Errorf("添加用户失败")
	}
	// 注册成功
	resp.Base.Code = base.Code_SUCCESS
	resp.Base.Msg = "注册成功"
	resp.UserId = uuid.String()
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = user.NewLoginResponse()
	resp.SetBase(base.NewBaseResp())
	// 查询用户
	u, err := s.UserRepo.QueryUserByAccount(req.Username, req.UserCategory)
	if err != nil {
		resp.Base.Code = base.Code_NOT_FOUND
		resp.Base.Msg = "用户不存在"
		return
	}
	// 验证密码是否正确
	if !ValidatePassword(req.Password, u.Salt, u.Password) {
		resp.Base.Code = base.Code_UNAUTHORIZED
		resp.Base.Msg = "用户密码错误"
		err = fmt.Errorf("invalid password")
		return
	}
	// 验证通过，登录成功
	resp.Base.SetCode(base.Code_SUCCESS)
	resp.Base.SetMsg("登录成功")
	resp.SetUserId(u.UserId)
	return
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, req *user.InfoRequest) (resp *user.InfoResponse, err error) {
	resp = user.NewInfoResponse()
	resp.SetBase(base.NewBaseResp())
	resp.SetUser(user.NewUser())

	usr, err := s.UserRepo.QueryUserById(req.UserId)
	if err != nil {
		resp.Base.Code = base.Code_NOT_FOUND
		resp.Base.Msg = "无用户信息"
		return
	}

	resp.Base.Code = base.Code_SUCCESS
	resp.Base.Msg = "成功查询用户信息"
	resp.User = pack.User(usr)

	return
}
