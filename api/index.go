package api

import (
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/router"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := gin.Default()
	router.SetApiRouter(server)
	router.SetWebRouter(server)
	server.ServeHTTP(w, r)
}
