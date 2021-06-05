package initRouter

import "GinHello/handler/user"

func UserRouter() {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", user.Register)
		userGroup.POST("/login", user.CreateJwt)
	}
}
