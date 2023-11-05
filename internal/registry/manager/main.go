package manager

import (
	"github.com/itihey/tikuAdapter/configs"
	"github.com/itihey/tikuAdapter/internal/registry"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"github.com/itihey/tikuAdapter/pkg/ratelimit"
	"gorm.io/gorm"
)

var defaultManager Manager

// Manager 所有的组件都注册到这里
type Manager struct {
	db        *gorm.DB
	config    configs.Config
	ipLimiter *ratelimit.IPRateLimiter
}

// RegistryManagerInterface manager interface
type RegistryManagerInterface interface {
	CloseManager() error
	GetDB() *gorm.DB
	GetIPLimiter() *ratelimit.IPRateLimiter
	GetConfig() configs.Config
}

// GetManager get manager
func GetManager() Manager {
	return defaultManager
}

// CreateManager create manager
func CreateManager() Manager {
	config := registry.Config()
	logger.SetupGinLog()
	db := registry.RegisterDB()
	defaultManager = Manager{
		db:        db,
		config:    config,
		ipLimiter: registry.Limit(config),
	}
	return defaultManager
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
