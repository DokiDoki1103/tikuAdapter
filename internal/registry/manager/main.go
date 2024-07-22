package manager

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/itihey/tikuAdapter/configs"
	"github.com/itihey/tikuAdapter/internal/registry"
	"github.com/itihey/tikuAdapter/internal/service"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"github.com/itihey/tikuAdapter/pkg/ratelimit"
	"gorm.io/gorm"
)

var defaultManager *Manager

// Manager 所有的组件都注册到这里
type Manager struct {
	db        *gorm.DB
	config    configs.Config
	ipLimiter *ratelimit.IPRateLimiter
	es        service.Elasticsearch
	bucket    *oss.Bucket
}

// RegistryManagerInterface manager interface
type RegistryManagerInterface interface {
	CloseManager() error
	GetDB() *gorm.DB
	GetIPLimiter() *ratelimit.IPRateLimiter
	GetConfig() configs.Config
	GetEs() *service.Elasticsearch
}

// GetManager get manager
func GetManager() *Manager {
	return defaultManager
}

// CreateManager create manager
func CreateManager() (*Manager, error) {
	config := registry.Config()
	logger.SetupGinLog()
	db := registry.RegisterDB(config)

	// 注册oss
	client, err := oss.New(config.OSS.EndPoint, config.OSS.AccessKeyID, config.OSS.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(config.OSS.BucketName)
	if err != nil {
		return nil, err
	}
	defaultManager = &Manager{
		db:        db,
		config:    config,
		ipLimiter: registry.Limit(config),
		es:        registry.RegisterEs(config),
		bucket:    bucket,
	}
	return defaultManager, nil
}

// CloseManager close manager
func (m Manager) CloseManager() error {
	return registry.CloseDB()
}

// GetDB get db
func (m Manager) GetDB() *gorm.DB {
	return m.db
}

// GetIPLimiter get ip limiter
func (m Manager) GetIPLimiter() *ratelimit.IPRateLimiter {
	return m.ipLimiter
}

// GetConfig get config
func (m Manager) GetConfig() configs.Config {
	return m.config
}

// GetEs get es
func (m Manager) GetEs() service.Elasticsearch {
	return m.es
}

// GetBucket get bucket
func (m Manager) GetBucket() *oss.Bucket {
	return m.bucket
}
