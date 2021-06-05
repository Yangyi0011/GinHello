package initRouter

import (
	"GinHello/handler/article"
)

// ArticleRouter 文章路由
func ArticleRouter() {
	articleGroup := router.Group("/article")
	{
		articleGroup.GET("/:id", article.GetOne)
		articleGroup.POST("/", article.Insert)
		articleGroup.GET("/", article.GetAll)
		articleGroup.DELETE("/:id", article.DeleteOne)
		articleGroup.PUT("/", article.Update)
	}
}
