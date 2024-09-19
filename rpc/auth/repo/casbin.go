package repo

import (
	"cloud_tinamic/pkg/pg"
	"cloud_tinamic/pkg/pgxadapter"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	SCHEMA = "auth"
	TABLE  = "router_permission"
)

var (
	modelPath string
	once      sync.Once
)

func initModelPath() {
	cwd, err := os.Getwd()
	if err != nil {
		klog.Fatalf("Failed to get current working directory: %v", err)
	}
	modelPath = filepath.Join(cwd, "repo", "model.conf")
}

func NewAdapter(pool *pg.PgPool) (*pgxadapter.Adapter, error) {
	return pgxadapter.NewAdapter("",
		pgxadapter.WithSchema(SCHEMA),
		pgxadapter.WithTableName(TABLE),
		pgxadapter.WithTimeout(1*time.Minute),
		pgxadapter.WithConnectionPool(pool.Pool))
}

func NewEnforcer(pool *pg.PgPool) *casbin.Enforcer {
	once.Do(initModelPath)

	a, err := NewAdapter(pool)
	if err != nil {
		klog.Fatalf("Failed to create adapter: %v", err)
	}

	e, err := casbin.NewEnforcer(modelPath, a)
	if err != nil {
		klog.Errorf("Failed to create enforcer: %v", err)
		return nil
	}

	if err := e.LoadPolicy(); err != nil {
		klog.Errorf("Failed to load policy: %v", err)
		return nil
	}

	return e
}
