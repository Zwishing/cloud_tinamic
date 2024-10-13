package repo

import (
	conf "cloud_tinamic/config"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
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
		var err error
		// 创建自定义 Logger
		customLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 日志输出位置
			logger.Config{
				LogLevel: logger.Info, // 日志级别
			},
		)
		DB, err = gorm.Open(postgres.Open(constr), &gorm.Config{
			Logger: customLogger,
		})
		if err != nil {
			klog.Errorf("failed to connect to database: %v", err)
			return
		}
		klog.Infof("successfully connected to db @ %s", cfg.GetString("database.postgresql.host"))
	})
	return DB
}
