package initRouter

import (
	"GinHello/initRouter"
	"github.com/gin-gonic/gin"
)

/*
	此文件定义有测试所需的资源，在测试时需要一起携带该文件进行测试
	测试命令：
		1、整个 test 包测试：
			在 test 目录下执行下面命令：
				go test		只有失败时才会打印日志
				go test -v  无论成功失败，都打印日志
		2、指定文件测试
			在 test 目录下执行下面命令：
				go test -v testInit.go indexGroupRouter_test.go
				只测试 indexGroupRouter_test.go 文件下的所有方法
		3、指定方法测试
			在 test 目录下执行下面命令：
				go test -v testInit.go indexGroupRouter_test.go -test.run TestUserRegister
				只会测试 indexGroupRouter_test.go 文件中的 TestUserRegister 方法
*/

// 全局变量
var (
	router *gin.Engine
)

// 初始化测试环境
func init() {
	// 设置当前文件的启动模式：TestMode
	gin.SetMode(gin.TestMode)
	router = initRouter.InitRouter()
	// 加载指定目录下的所有模板文件
	if mode := gin.Mode(); mode == gin.TestMode {
		// 测试模式下要使用绝对路径，否则在单元测试中会找不到目标
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
}

