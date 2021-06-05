package initRouter

import (
	"GinHello/handler"
	"GinHello/middleware"
)

// UserGroupRouter 用户资源处理路由
func UserGroupRouter() {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/:name", handler.UserSave)
		userGroup.GET("", handler.UserSaveByQuery)
		userGroup.POST("/register", handler.UserRegister)
		userGroup.POST("/login", handler.UserLogin)
		// middleware.Auth() 是授权认证中间件，如此用法表示局部使用中间件
		userGroup.GET("/profile", middleware.Auth(), handler.UserProfile)
		userGroup.POST("/update", middleware.Auth(), handler.UpdateUserProfile)
	}
}
