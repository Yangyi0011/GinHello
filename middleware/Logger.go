package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Logger 自定义日志中间件
// 中间件需要返回 gin.HandlerFunc 函数，所以定义返回函数
// 中间件有个 Next 函数，在我们定义的众多中间件，会形成一条中间件链，
//	而通过 Next 函数来对后面的中间件进行执行。
// Next 函数是在请求前执行，而 Next 函数后是在请求后执行。
func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		host := context.Request.Host
		url := context.Request.URL
		method := context.Request.Method
		fmt.Printf("%s::%s \t %s \t %s ", time.Now().Format("2006-01-02 15:04:05"), host, url, method)
		// Next() 之前的代码是在请求之前执行
		context.Next()
		// Next() 之后的代码是在请求之后执行
		fmt.Println(context.Writer.Status())
	}
}

// Auth 局部授权认证中间件
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		cookie, e := context.Request.Cookie("user_cookie")
		if e != nil {
			// 表示对当前的请求进行中止，确保剩余的处理程序对该请求不再处理
			context.Abort()
			// 响应到 401 页面
			context.HTML(http.StatusUnauthorized, "401.tmpl", nil)
			return
		}
		println("已授权")
		// 认证成功时刷新 cookie 的过期时间
		context.SetCookie(cookie.Name, cookie.Value, 1000, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		context.Next()
	}
}