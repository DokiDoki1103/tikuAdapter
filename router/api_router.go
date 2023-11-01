package router

import (
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/controller"
)

// SetAPIRouter 设置API路由
func SetAPIRouter(router *gin.Engine) {
	apiRouter := router.Group("/adapter-service")

	apiRouter.POST("/search", controller.Search)
	apiRouter.POST("/parser", controller.Parse)
}
