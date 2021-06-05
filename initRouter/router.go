package initRouter

import (
	"GinHello/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/*
	此文件已不再使用，是之前学习时的记录
*/
// SetupRouter 设置启动路由
func SetupRouter() *gin.Engine {
	// 添加 Get 请求路由
	router.GET("/", retHelloGinAndMethod)
	// 添加 Post 请求路由
	router.POST("/", retHelloGinAndMethod)
	// 添加 Put 请求路由
	router.PUT("/", retHelloGinAndMethod)
	// 添加 Delete 请求路由
	router.DELETE("/", retHelloGinAndMethod)
	// 添加 Patch 请求路由
	router.PATCH("/", retHelloGinAndMethod)
	// 添加 Head 请求路由
	router.HEAD("/", retHelloGinAndMethod)
	// 添加 Options 请求路由
	router.OPTIONS("/", retHelloGinAndMethod)

	// 添加带请求参数的 Get 请求路由，/user/:name 相当于 /user/{name}
	router.GET("/user/:name", handler.UserSave)

	// 通过 context.Query() 获取参数
	router.GET("/user", handler.UserSaveByQuery)
	return router
}

// GroupRouter 路由分组
func GroupRouter() *gin.Engine {
	index := router.Group("/")
	{
		// index 组内的所有路由都指向同一路径 "/"
		// 添加 Get 请求路由
		index.GET("", retHelloGinAndMethod)
		// 添加 Post 请求路由
		index.POST("", retHelloGinAndMethod)
		// 添加 Put 请求路由
		index.PUT("", retHelloGinAndMethod)
		// 添加 Delete 请求路由
		index.DELETE("", retHelloGinAndMethod)
		// 添加 Patch 请求路由
		index.PATCH("", retHelloGinAndMethod)
		// 添加 Head 请求路由
		index.HEAD("", retHelloGinAndMethod)
		// 添加 Options 请求路由
		index.OPTIONS("", retHelloGinAndMethod)
	}

	// 使用 Any 请求方法来代替上面的各种方法
	index.Any("", retHelloGinAndMethod)

	user := router.Group("/user")
	{
		user.GET("/:name", handler.UserSave)
		user.GET("", handler.UserSaveByQuery)
	}

	return router
}

// AnyRouter 使用 Any 请求方法来代替 index 所支持的一堆方法
func AnyRouter() *gin.Engine {
	index := router.Group("/")
	// 使用 Any 请求方法来代替 index 所支持的各种方法
	index.Any("", retHelloGinAndMethod)
	return router
}

// BaseRouter 基础路由
func BaseRouter() *gin.Engine {
	index := router.Group("/")
	index.Any("", handler.Index)

	// 添加 user
	userGroup := router.Group("/user")
	{
		userGroup.GET("/:name", handler.UserSave)
		userGroup.GET("", handler.UserSaveByQuery)
		userGroup.POST("/register", handler.UserRegister)
	}

	return router
}

// 抽取响应体
func retHelloGinAndMethod(context *gin.Context) {
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
}