package configs

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/itihey/tikuAdapter/internal/search"
)

// Config 所有的配置文件
type Config struct {
	Limit             LimitConfig          `yaml:"limit"`
	API               []search.API         `yaml:"api"`
	Elasticsearch     elasticsearch.Config `yaml:"elasticsearch"`
	RecordEmptyAnswer bool                 `yaml:"recordEmptyAnswer"`
	Mysql             string               `yaml:"mysql"`
}

// LimitConfig 限流配置
type LimitConfig struct {
	Enable   bool   `yaml:"enable"`
	Duration uint   `yaml:"duration"`
	Requests uint64 `yaml:"requests"`
}
