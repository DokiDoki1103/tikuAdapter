package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/api"
	"github.com/itihey/tikuAdapter/internal/registry/manager"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

//go:embed dist
var buildFS embed.FS

//go:embed dist/index.html
var indexPage []byte

func main() {
	mg := manager.CreateManager()
	defer func(mg manager.Manager) {
		err := mg.CloseManager()
		if err != nil {
			logger.FatalLog(err)
		}
	}(mg)

	quitSignal := make(chan os.Signal, 1)
	signal.Notify(
		quitSignal,
		syscall.SIGINT, syscall.SIGTERM,
	)
	go func() {
		server := gin.Default()
		api.SetAPIRouter(server)
		api.SetWebRouter(buildFS, indexPage, server)
		err := server.Run("0.0.0.0:8060")
		if err != nil {
			logger.FatalLog(err)
		}
	}()
	<-quitSignal
}
