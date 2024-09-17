package repo

import (
	conf "cloud_tinamic/config"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	DB     *gorm.DB
	dbOnce sync.Once
)

func NewDB() *gorm.DB {
	dbOnce.Do(func() {
		cfg := conf.GetConfigInstance()
		constr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
			cfg.GetString("database.postgresql.user"),
			cfg.GetString("database.postgresql.password"),
			cfg.GetString("database.postgresql.host"),
			cfg.GetInt32("database.postgresql.port"),
			cfg.GetString("database.postgresql.database"),
			cfg.GetString("database.postgresql.sslmode"))
		DB, _ = gorm.Open(postgres.Open(constr), &gorm.Config{})
		klog.Infof("success connect db @ %s", cfg.GetString("database.postgresql.host"))
	})
	return DB
}
