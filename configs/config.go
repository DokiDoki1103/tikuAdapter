package configs

import "github.com/itihey/tikuAdapter/internal/search"

// Config 所有的配置文件
type Config struct {
	Limit LimitConfig  `yaml:"limit"`
	API   []search.API `yaml:"api"`
}

// LimitConfig 限流配置
type LimitConfig struct {
	Enable   bool   `yaml:"enable"`
	Duration uint   `yaml:"duration"`
	Requests uint64 `yaml:"requests"`
}
