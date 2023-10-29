package api

import (
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/router"
	"net/http"
)

// Handler vercel部署需要
func Handler(w http.ResponseWriter, r *http.Request) {
	server := gin.Default()
	router.SetAPIRouter(server)
	router.SetWebRouter(server)
	server.ServeHTTP(w, r)
}
