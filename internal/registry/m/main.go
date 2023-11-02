package m

import (
	"github.com/itihey/tikuAdapter/configs"
	"github.com/itihey/tikuAdapter/internal/dao"
	"github.com/itihey/tikuAdapter/internal/registry"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"github.com/itihey/tikuAdapter/pkg/ratelimit"
)

var defaultManager Manager

// Manager db manager
type Manager struct {
	Config    *configs.Config
	Query     *dao.Query
	IPLimiter *ratelimit.IPRateLimiter
}

// GetManager get db manager
func GetManager() Manager {
	return defaultManager
}

// CreateManager create db manager
func CreateManager() {
	config := registry.Config()
	logger.SetupGinLog()
	defaultManager = Manager{
		Config:    config,
		Query:     registry.DB(),
		IPLimiter: registry.Limit(config),
	}
}
