package main

import (
	"github.com/itihey/tikuAdapter/api"
	"github.com/itihey/tikuAdapter/internal/registry/s"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"net/http"
)

func main() {
	s.CreateManager()

	logger.SetupGinLog()
	http.HandleFunc("/", api.Handler)
	if err := http.ListenAndServe(":8060", nil); err != nil {
		logger.FatalLog(err)
	}
}
