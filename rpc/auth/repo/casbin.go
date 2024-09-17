package repo

import (
	"cloud_tinamic/pkg/pg"
	"cloud_tinamic/pkg/pgxadapter"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"
	"time"
)

const (
	SCHEMA = "auth"
	TABLE  = "router_permission"
)

var modelPath string

func init() {
	// 获取当前工作目录
	cwd, err := os.Getwd()
	if err != nil {
		klog.Fatalf("Failed to create current dir: %v", err)
	}
	modelPath = fmt.Sprintf("%s%s", cwd, "/rpc/auth/repo/model.conf")
}

func NewAdapter(pool *pg.PgPool) (*pgxadapter.Adapter, error) {
	a, err := pgxadapter.NewAdapter("",
		pgxadapter.WithSchema(SCHEMA),
		pgxadapter.WithTableName(TABLE),
		pgxadapter.WithTimeout(1*time.Minute),
		pgxadapter.WithConnectionPool(pool.Pool))
	return a, err
}

func NewEnforcer(pool *pg.PgPool) *casbin.Enforcer {
	a, err := NewAdapter(pool)
	if err != nil {
		klog.Fatalf("Failed to create adapter: %v", err)
	}

	e, err := casbin.NewEnforcer(modelPath, a)
	if err != nil {
		klog.Fatalf("Failed to create enforcer: %v", err)
		return nil
	}

	err = e.InitWithAdapter(modelPath, a)
	if err != nil {
		return nil
	}

	return e
}
