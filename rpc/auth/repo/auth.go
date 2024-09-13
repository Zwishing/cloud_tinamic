package repo

import (
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/kitex/pkg/klog"
)

type AuthRepo interface {
	Auth(sub, obj, act string) bool
}

type AuthRepoImpl struct {
	*casbin.Enforcer
}

func NewAuthRepoImpl(e *casbin.Enforcer) AuthRepo {
	return &AuthRepoImpl{
		e,
	}
}

func (e *AuthRepoImpl) Auth(sub, obj, act string) bool {
	// 使用 EnforceEx 方法进行权限校验
	ok, _, err := e.EnforceEx(sub, obj, act)
	if err != nil {
		klog.Errorf("Failed to enforce policy: %v", err)
		return false
	}
	return ok
}

func (e *AuthRepoImpl) AddPolicys(sub, obj, act string) (bool, error) {
	//policy, err := e.AddPolicy(sub, obj, act)
	//if err != nil {
	//	return false, err
	//}
	return false, nil
}
