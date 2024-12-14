package api

import (
	"embed"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/controller"
	"github.com/itihey/tikuAdapter/internal/middleware"
	"io/fs"
	"net/http"
)

// SetAPIRouter 设置API路由
func SetAPIRouter(router *gin.Engine) {

	router.Any("/sqp/*path", controller.Proxy)

	apiRouter := router.Group("/adapter-service")

	apiRouter.Use(middleware.GlobalAPIRateLimit) // 全局限流

	apiRouter.POST("/upload", controller.UploadFile)
	apiRouter.GET("/plat", controller.Plat)

	apiRouter.POST("/search", controller.Search)
	apiRouter.POST("/parser", controller.Parser)

	apiRouter.GET("/courses", controller.Auth, controller.Courses)
	apiRouter.POST("/questions/search", controller.Auth, controller.GetQuestions)
	apiRouter.POST("/questions", controller.Auth, controller.CreateQuestion)
	apiRouter.PUT("/questions/:id", controller.Auth, controller.UpdateQuestions)
	apiRouter.DELETE("/questions/:id", controller.Auth, controller.DeleteQuestion)

	apiRouter.GET("/user/login", controller.UserLogin)                    // 登录接口
	apiRouter.GET("/user", controller.Auth, controller.UserList)          // 自己创建的用户列表
	apiRouter.POST("/user", controller.Auth, controller.CreateUser)       // 创建用户
	apiRouter.DELETE("/user/:id", controller.Auth, controller.DeleteUser) // 删除用户

	apiRouter.GET("/logs", controller.Auth, controller.LogList) // 日志列表

}

type embedFileSystem struct {
	http.FileSystem
}

// Exists 判断文件是否存在
func (e embedFileSystem) Exists(_ string, path string) bool {
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

// SetWebRouter 设置web路由
func SetWebRouter(buildFS embed.FS, indexPage []byte, router *gin.Engine) {
	router.Use(static.Serve("/", embedFolder(buildFS, "dist")))
	router.NoRoute(func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexPage)
	})
}
