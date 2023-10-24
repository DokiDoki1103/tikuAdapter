package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/api/router"
	"log"
)

//go:embed dist
var buildFS embed.FS

//go:embed dist/index.html
var indexPage []byte

// main -
func main() {
	server := gin.Default()
	router.SetRouter(server, buildFS, indexPage)

	err := server.Run("0.0.0.0:8060")
	if err != nil {
		log.Println(err)
	}
}
