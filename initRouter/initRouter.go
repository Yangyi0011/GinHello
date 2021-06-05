package initRouter

import (
	"github.com/gin-gonic/gin"
)

var (
	// 全局路由引擎
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	SwaggerRouter()
	ArticleRouter()
	return router
}