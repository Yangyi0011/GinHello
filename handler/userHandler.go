package handler

import (
	"GinHello/model"
	"GinHello/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// UserSave 保存 User
func UserSave(context *gin.Context) {
	// 通过 Param() 获取路由路径中的请求参数
	username := context.Param("name")
	context.String(http.StatusOK, "用户【"+username+"】已保存")
}

// UserSaveByQuery 通过 Query() 获取参数
func UserSaveByQuery(context *gin.Context) {
	username := context.DefaultQuery("name", "未知")
	age := context.DefaultQuery("age", "0")
	context.String(http.StatusOK, "用户："+username+"，年龄："+age+"已经保存")
}

// UserRegister 用户注册
func UserRegister(context *gin.Context) {
	// context = {*github.com/gin-gonic/gin.Context | 0xc00037e100} 方法一：手动接收表单信息
	//// 使用 PostForm 提取 From 表单指定的 key 对应的值，没有对应的 key 则返回 “”
	//email := context.PostForm("email")
	//// 使用 DefaultPostForm 提取 From 表单指定的 key 对应的值，可以设置默认值
	//password := context.DefaultPostForm("password", "123456")
	//rePassword := context.DefaultPostForm("rePassword", "123456")
	//println("email=", email, ", password=", password, ", rePassword=", rePassword)

	// 方法二：使用表单绑定的方式来自动接收表单信息
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		// 使用 log 来打印日志
		log.Println("UserRegister error->", err.Error())
		// 响应错误到 error 页面
		context.HTML(http.StatusOK, "error.tmpl", gin.H {
			"error": "输入的数据不合法",
		})
		return
	}
	// 保存用户注册信息
	id := user.Save()
	log.Println("save user id = ", id)
	// 重定向到首页
	context.Redirect(http.StatusMovedPermanently, "/")
}

// UserLogin 用户登录
func UserLogin(context *gin.Context) {
	var user model.UserModel
	if err := context.Bind(&user); err != nil {
		log.Println("login 数据绑定错误", err.Error())
		context.HTML(http.StatusOK, "error.tmpl", gin.H {
			"error": err,
		})
		return
	}
	// 通过 email 查询用户信息
	u := user.QueryByEmail()
	if u.Password != user.Password {
		log.Println("用户名或密码输入错误")
		context.HTML(http.StatusOK, "error.tmpl", gin.H {
			"error": "用户名或密码输入错误",
		})
		return
	}
	log.Println("登录成功", u.Email)

	// 设置 cookie
	// name：名称，value：值，maxAge：有效时间，path：所在的目录，domain：所在域，secure：是否只能通过 https 访问，httpOnly：是否可以通过 js 代码进行操作
	context.SetCookie("user_cookie", strconv.Itoa(u.Id), 1000, "/", "localhost", false, true)
	// 响应到前端页面，并携带数据 email 过去
	context.HTML(http.StatusOK, "index.tmpl", gin.H {
		"email": u.Email,
		"id":    u.Id,
	})
}

// UserProfile 获取用户信息
func UserProfile(context *gin.Context) {
	id := context.Query("id")
	var user model.UserModel
	i, err := strconv.Atoi(id)
	u, e := user.QueryById(i)
	if e != nil || err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H {
			"error": e,
		})
		return
	}
	context.HTML(http.StatusOK, "user_profile.tmpl", gin.H {
		"user": u,
	})
}

// UpdateUserProfile 上传用户头像
func UpdateUserProfile(context *gin.Context) {
	var user model.UserModel
	// 通过 context.ShouldBind() 绑定其他数据
	if err := context.ShouldBind(&user); err != nil {
		log.Println("数据绑定发生错误 ", err.Error())
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}
	// 通过 context.FormFile() 将文件数据进行获取。
	file, e := context.FormFile("avatar-file")
	if e != nil {
		log.Println("文件上传错误", e.Error())
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		return
	}
	// 获取当前项目的根路径
	path := utils.RootPath()
	// 组装文件上传目录
	path = filepath.Join(path, "avatar")
	log.Println("path =>", path)
	// 判断目标路径是否存在
	exist, e := utils.PathIsExist(path)
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("创建文件夹错误", e.Error())
	}
	// 目录不存在才创建
	if !exist {
		e = os.MkdirAll(path, os.ModePerm)
		if e != nil {
			context.HTML(http.StatusOK, "error.tmpl", gin.H{
				"error": e,
			})
			log.Panicln("无法创建文件夹", e.Error())
		}
	}
	// 给上传文件重命名
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	// 保存上传文件
	e = context.SaveUploadedFile(file, filepath.Join(path, fileName))
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法保存文件", e.Error())
	}
	avatarUrl := "/avatar/" + fileName
	user.Avatar = sql.NullString{String: avatarUrl}
	e = user.Update(user.Id)
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("数据无法更新", e.Error())
	}
	// 使用重定向来刷新页面内容
	context.Redirect(http.StatusMovedPermanently, "/user/profile?id="+strconv.Itoa(user.Id))
}