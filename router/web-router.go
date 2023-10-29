package router

import (
	"embed"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

type embedFileSystem struct {
	http.FileSystem
}

// Exists 判断文件是否存在
func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func embedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	efs, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(efs),
	}
}

//go:embed dist
var buildFS embed.FS

//go:embed dist/index.html
var indexPage []byte

// SetWebRouter 设置web路由
func SetWebRouter(router *gin.Engine) {
	router.Use(static.Serve("/", embedFolder(buildFS, "router/dist")))
	router.NoRoute(func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexPage)
	})
}
