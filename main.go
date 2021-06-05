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

// @title Gin swagger
// @version 1.0
// @description Gin swagger 示例项目

// @contact.name yangyi
// @contact.email 1024569696@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	// 启动
	err := router.Run()
	if err != nil {
		log.Fatal("项目启动错误：", err)
		return
	}
}