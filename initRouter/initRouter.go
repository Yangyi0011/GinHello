package initRouter

import (
	"GinHello/middleware"
	"GinHello/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

var (
	// 全局路由引擎
	router *gin.Engine
)

func init() {
	//router = gin.Default()

	// 不再使用默认的 gin 引擎，使用我们自己配置的
	router = gin.New()
	// 全局使用自定义中间件
	router.Use(middleware.Logger(), gin.Recovery())
}

// StaticsRouter 静态资源路由
func StaticsRouter() {
	// 加载指定目录下的所有模板文件
	if mode := gin.Mode(); mode == gin.TestMode {
		// 测试模式下要使用绝对路径，否则在单元测试中会找不到目标
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	// 加载静态资源
	router.Static("/statics", "./statics")
	// 设置 icon，需要将 .ico 文件放在项目的根目录下
	router.StaticFile("/favicon.ico","./favicon.ico")
	// 设置用户头像存储路径
	router.StaticFS("/avatar", http.Dir(filepath.Join(utils.RootPath(), "avatar")))
}

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	SwaggerRouter()
	IndexGroupRouter()
	UserGroupRouter()
	StaticsRouter()
	ArticleRouter()
	return router
}