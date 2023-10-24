package router

import (
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/api/controller"
)

func SetApiRouter(router *gin.Engine) {
	apiRouter := router.Group("/adapter-service")

	apiRouter.POST("/search", controller.Search)
}
