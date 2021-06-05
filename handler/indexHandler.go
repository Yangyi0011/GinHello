package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Index 首页
func Index(context *gin.Context) {
	// 被请求后将响应一个 HTML 页面
	context.HTML(http.StatusOK, "index.tmpl", gin.H {
		// 把遍历 title 渲染到 index.tmpl 文件中，并向请求响应一个 HTML 页面内容
		"title": "hello gin " + strings.ToLower(context.Request.Method) + " method",
	})
}
