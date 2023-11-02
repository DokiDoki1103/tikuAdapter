package registry

import (
	"github.com/itihey/tikuAdapter/configs"
	"github.com/itihey/tikuAdapter/pkg/ratelimit"
	"golang.org/x/time/rate"
)

// Limit get ratelimit instance
func Limit(cfg *configs.Config) *ratelimit.IPRateLimiter {
	limit := cfg.Limit
	r := rate.Limit(float64(limit.LimitDuration) / float64(limit.LimitRequests))
	return ratelimit.NewIPRateLimiter(r, 1)
}
