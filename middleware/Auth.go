package middleware

import (
	"GinHello/config"
	"GinHello/model"
	"GinHello/utils/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Auth 登录认证
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		result := model.Result{}
		// 首先在请求头获取 token
		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.Abort()
			result.Code = http.StatusUnauthorized
			result.Message = "请先登录"
			response.Fail(context, result)
			return
		}
		// 然后对先把 token 进行解析，将 Bearer 和 JWT 拆分出来，将 JWT 进行校验。
		auth = strings.Fields(auth)[1]
		// 校验token
		_, err := parseToken(auth)
		if err != nil {
			context.Abort()
			result.Code = http.StatusUnauthorized
			result.Message = "token 过期：" + err.Error()
			response.Fail(context, result)
			return
		}
		println("token 正确")
		context.Next()
	}
}

// 解析 token
func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(config.Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
