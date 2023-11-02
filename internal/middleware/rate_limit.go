package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/registry/s"
	"net/http"
)

// GlobalAPIRateLimit 全局API限流
func GlobalAPIRateLimit(c *gin.Context) {
	manager := s.GetManager()
	if manager.Config.Limit.Enable && !manager.IPLimiter.GetLimiter(c.ClientIP()).Allow() {
		c.AbortWithStatus(http.StatusTooManyRequests)
	}
}
