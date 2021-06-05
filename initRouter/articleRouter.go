package initRouter

import (
	"GinHello/handler/article"
	"GinHello/middleware"
)

// ArticleRouter 文章路由
func ArticleRouter() {
	articleGroup := router.Group("/article")
	{
		articleGroup.GET("/:id", article.GetOne)
		articleGroup.POST("/", article.Insert)
		articleGroup.GET("/", middleware.Auth(), article.GetAll)
		articleGroup.DELETE("/:id", article.DeleteOne)
		articleGroup.PUT("/", article.Update)
	}
}
