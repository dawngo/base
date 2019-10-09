package routers

import (
	"github.com/Brave-man/base/controllers/baseController"
	"github.com/Brave-man/base/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.New()

	// 使用文件log写入器中间件
	router.Use(middleware.LoggerToFile())
	router.Use(gin.Recovery())

	// --------------------路由表----------------------
	router.GET("/hello", baseController.GetIndex)
	router.POST("/hello", baseController.PostIndex)

	return router
}
