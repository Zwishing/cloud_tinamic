package conf

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
	"sync"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	*viper.Viper
}

func NewConfig() *Config {
	config := &Config{Viper: viper.New()}

	config.AddConfigPath("./config")
	config.AddConfigPath("../config")
	config.AddConfigPath("../../config")
	config.SetConfigName("tinamic")
	config.SetConfigType("toml")

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			klog.Fatalf("无法读取配置:%v", err)
		}
	}

	return config
}

func GetConfigInstance() *Config {
	once.Do(func() {
		cfg = NewConfig()
	})
	return cfg
}
