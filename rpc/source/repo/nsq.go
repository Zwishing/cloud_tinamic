package repo

type NsqConfig struct {
	Topic       string
	Channel     string
	Address     string
	MaxAttempts uint16
}

func NewNsqConfig() *NsqConfig {
	// TODO 从配置文件读取
	return &NsqConfig{
		Topic:       "add_data",
		Channel:     "channel1",
		Address:     "1.92.113.25:4161",
		MaxAttempts: 10,
	}
}
