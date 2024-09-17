package repo

import (
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/kitex/pkg/klog"
)

type AuthRepo interface {
	Auth(sub, obj, act string) bool
	AddPolicy(sub, obj, act string) bool
}

type AuthRepoImpl struct {
	enforcer *casbin.Enforcer
}

func NewAuthRepoImpl(e *casbin.Enforcer) AuthRepo {
	return &AuthRepoImpl{
		e,
	}
}

func (e *AuthRepoImpl) Auth(sub, obj, act string) bool {
	// 使用 EnforceEx 方法进行权限校验
	ok, _, err := e.enforcer.EnforceEx(sub, obj, act)
	if err != nil {
		klog.Errorf("Failed to enforce policy: %v", err)
		return false
	}
	return ok
}

func (e *AuthRepoImpl) AddPolicy(sub, obj, act string) bool {
	policy, err := e.enforcer.AddPolicy(sub, obj, act)
	if err != nil {
		klog.Errorf("Failed to add policy: %v", err)
		return false
	}
	// If the rule already exists, the function returns false and the rule will not be added.
	if policy == false {
		klog.Infof("The rule already exists: %s %s %s", sub, obj, act)
		return true
	}
	return policy
}
