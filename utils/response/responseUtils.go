package response

import (
	"GinHello/model"
	"github.com/gin-gonic/gin"
)

const (
	successCode = 200
	failCode = -99
)

// Success 响应成功
func Success(context *gin.Context, result model.Result) {
	if result.Code == 0 {
		result.Code = successCode
	}
	if result.Message == "" {
		result.Message = "成功"
	}
	context.JSON(result.Code, gin.H{
		"result": result,
	})
}

// Fail 响应失败
func Fail(context *gin.Context, result model.Result) {
	if result.Code == 0 {
		result.Code = failCode
	}
	if result.Message == "" {
		result.Message = "失败"
	}
	context.JSON(result.Code, gin.H{
		"result": result,
	})
}