package conf

import (
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
	"os"
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
	config := new(Config)
	config.Viper = viper.New()

	config.AddConfigPath("./config")
	config.AddConfigPath("../config")
	config.AddConfigPath("../../config")
	config.SetConfigName("tinamic")
	config.SetConfigType("toml")

	// Read configuration
	if err := config.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			klog.Errorf("failed to read configuration:%v", err)
			os.Exit(1)
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
