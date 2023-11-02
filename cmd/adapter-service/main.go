package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/api"
	"github.com/itihey/tikuAdapter/internal/registry/m"
	"github.com/itihey/tikuAdapter/pkg/logger"
)

//go:embed dist
var buildFS embed.FS

//go:embed dist/index.html
var indexPage []byte

func main() {
	m.CreateManager()

	server := gin.Default()
	api.SetAPIRouter(server)
	api.SetWebRouter(buildFS, indexPage, server)
	err := server.Run(":8060")
	if err != nil {
		logger.FatalLog(err)
	}
}
