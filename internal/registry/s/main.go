package s

import (
	"github.com/itihey/tikuAdapter/configs"
	"github.com/itihey/tikuAdapter/internal/dao"
	"github.com/itihey/tikuAdapter/internal/registry"
	"github.com/itihey/tikuAdapter/pkg/ratelimit"
)

var defaultManager Manager

type Manager struct {
	Config    *configs.Config
	Query     *dao.Query
	IPLimiter *ratelimit.IPRateLimiter
}

// GetManager get db manager
func GetManager() Manager {
	return defaultManager
}

func CreateManager() {
	config := registry.Config()
	defaultManager = Manager{
		Config:    config,
		Query:     registry.DB(),
		IPLimiter: registry.Limit(config),
	}
}
