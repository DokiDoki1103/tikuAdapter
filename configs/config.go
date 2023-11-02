package configs

// Config 所有的配置文件
type Config struct {
	Limit LimitConfig
}

type LimitConfig struct {
	Enable        bool
	LimitDuration uint
	LimitRequests uint64
}
