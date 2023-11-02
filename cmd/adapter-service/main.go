package main

import (
	"github.com/itihey/tikuAdapter/api"
	"github.com/itihey/tikuAdapter/internal/registry/m"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"net/http"
)

func main() {
	m.CreateManager()

	logger.SetupGinLog()
	http.HandleFunc("/", api.Handler)
	if err := http.ListenAndServe(":8060", nil); err != nil {
		logger.FatalLog(err)
	}
}
