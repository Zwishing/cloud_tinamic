package repo

import (
	"cloud_tinamic/pkg/pg"
	"cloud_tinamic/pkg/pgxadapter"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

const (
	SCHEMA = "user_info"
	TABLE  = "user_permission"
)

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
	e, err := casbin.NewEnforcer("model.conf", a)
	if err != nil {
		klog.Fatalf("Failed to create enforcer: %v", err)
		return nil
	}
	return e
}
