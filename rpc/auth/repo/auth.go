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
	return &AuthRepoImpl{enforcer: e}
}

func (e *AuthRepoImpl) Auth(sub, obj, act string) bool {
	ok, err := e.enforcer.Enforce(sub, obj, act)
	if err != nil {
		klog.Errorf("Failed to enforce policy: %v", err)
		return false
	}
	return ok
}

func (e *AuthRepoImpl) AddPolicy(sub, obj, act string) bool {
	success, err := e.enforcer.AddPolicy(sub, obj, act)
	if err != nil {
		klog.Errorf("Failed to add policy: %v", err)
		return false
	}
	if !success {
		klog.Infof("The rule already exists: %s %s %s", sub, obj, act)
	}
	return true
}
