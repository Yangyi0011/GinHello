package user

import (
	"GinHello/config"
	"GinHello/model"
	"GinHello/utils/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// CreateJwt 创建 JWT
// @Summary 用户登录
// @Description 用户登录签发 JWT
// @Tags 用户管理
// @version 1.0
// @Accept json
// @Produce  json
// @Param article body model.User true "用户"
// @Success 200 object model.Result 登录成功
// @Failure -99 object model.Result 登录失败
// @Router /user/login [post]
func CreateJwt(ctx *gin.Context) {
	user := &model.User{}
	result := model.Result{}
	// 数据绑定
	if e := ctx.BindJSON(&user); e != nil {
		result.Message = "数据绑定失败"
		response.Fail(ctx, result)
		return
	}
	// 通过 username 到数据库查询数据来对比
	u := user.QueryByUsername()
	if u.Password != user.Password {
		result.Message = "用户名或密码错误"
		response.Fail(ctx, result)
		return
	}

	// 签发 JWT
	expiresTime := time.Now().Unix() + int64(config.OneDayOfHours)
	claims := jwt.StandardClaims{
		Audience:  user.Username,              // 受众
		ExpiresAt: expiresTime,                // 失效时间
		Id: strconv.Itoa(int(user.ID)), 	   // 编号
		IssuedAt:  time.Now().Unix(),          // 签发时间
		Issuer:    "gin hello",                // 签发人
		NotBefore: time.Now().Unix(),          // 生效时间
		Subject:   "login",                    // 主题
	}

	// 加密JWT
	jwtSecret := []byte(config.Secret)
	// 通过 HS256 算法生成 tokenClaims ,这就是我们的 HEADER 部分和 PAYLOAD。
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		result.Message = "登录失败"
		response.Fail(ctx, result)
		return
	}
	result.Message = "登录成功"
	result.Data = "Bearer " + token
	result.Code = http.StatusOK
	response.Success(ctx, result)
}

// Register 用户注册
// @Summary 用户注册
// @Description 注册用户并保存到数据库
// @Tags 用户管理
// @version 1.0
// @Accept json
// @Produce  json
// @Param article body model.User true "用户"
// @Success 200 object model.Result 成功后返回值
// @Failure -99 object model.Result 注册失败
// @Router /user/register [post]
func Register(ctx *gin.Context) {
	user := model.User{}
	result := model.Result{}
	// 数据绑定
	if e := ctx.BindJSON(&user); e != nil {
		result.Message = "数据绑定失败"
		response.Fail(ctx, result)
		return
	}
	if !user.Insert() {
		result.Message = "注册失败"
		response.Fail(ctx, result)
		return
	}
	result.Message = "注册成功"
	result.Code = http.StatusOK
	result.Data = user.ID
	response.Success(ctx, result)
}