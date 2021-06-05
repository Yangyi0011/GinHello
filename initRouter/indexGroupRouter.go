package initRouter

import "GinHello/handler"

// IndexGroupRouter 根目录路由
func IndexGroupRouter() {
	index := router.Group("/")
	index.Any("", handler.Index)
}
