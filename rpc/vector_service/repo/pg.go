package repo

import (
	"cloud_tinamic/config"
	"cloud_tinamic/pkg/pg"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
)

var (
	db *pg.PgPool //定义一个连接池
)

func NewPgPool() *pg.PgPool {
	cfg := conf.GetConfigInstance()
	constr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.GetString("database.postgresql.user"),
		cfg.GetString("database.postgresql.password"),
		cfg.GetString("database.postgresql.host"),
		cfg.GetInt32("database.postgresql.port"),
		cfg.GetString("database.postgresql.database"),
		cfg.GetString("database.postgresql.sslmode"))
	dbConfig := pg.NewPgConfig(pg.WithConnString(constr))
	db, _ = pg.NewPgPool(dbConfig.Config)
	klog.Infof("success connect db @ %s", cfg.GetString("database.postgresql.host"))
	return db
}
