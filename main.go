package main

import (
	"GinHello/initRouter"
	"github.com/gin-gonic/gin"
	"log"
)

// 全局变量
var (
	router *gin.Engine
)

// 资源初始化
func init() {
	// 初始化路由
	router = initRouter.InitRouter()
}

func main() {
	// 启动
	err := router.Run()
	if err != nil {
		log.Fatal("项目启动错误：", err)
		return
	}
}