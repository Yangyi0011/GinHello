package initRouter

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestUserGroupRouter(t *testing.T) {
	// 测试 Restful 风格的 GET 请求
	username := "jerry"
	// 复用 w 会导致断言不通过
	w1 := httptest.NewRecorder()
	req2, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w1, req2)
	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, "用户【"+username+"】已保存", w1.Body.String())

	// 测试传统 GET 请求
	age := 20
	// 复用 w 会导致断言不通过
	w2 := httptest.NewRecorder()
	req3, _ := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+strconv.Itoa(age), nil)
	router.ServeHTTP(w2, req3)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, "用户："+username+"，年龄："+strconv.Itoa(age)+"已经保存", w2.Body.String())
}

// 测试用户表单注册
func TestUserRegister(t *testing.T) {
	//	要测试表单请求，首先我们要构造一个结构，该结构是为了帮助我们将我们要提交的信息存放到表单中，
	//	之后还要指定请求头信息。
	value := url.Values{}
	value.Add("email", "tom@qq.com")
	value.Add("password", "123456")
	value.Add("rePassword", "123456")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/register",
		bytes.NewBufferString(value.Encode()))
	// 指定请求头信息
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusMovedPermanently, w.Code)
}

// 测试模型的数据校验功能
func TestModelValidator(t *testing.T) {
	value := url.Values{}
	value.Add("email", "tom")
	value.Add("password", "123")
	value.Add("rePassword", "123456")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/register",
		bytes.NewBufferString(value.Encode()))
	// 指定请求头信息
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusMovedPermanently, w.Code)

	/*
		测试虽然通过了，但会有两行错误信息：
			UserRegister error-> Key: 'UserModel.Email' Error:Field validation for 'Email' failed on the 'email tag
				Key: 'UserModel.RePassword' Error:Field validation for 'RePassword' failed on the 'eqfield' tag
		该信息说明了我们的 Email 和 RePassword 信息校验没有通过。
	*/
}

// 测试用户登录
func TestUserLogin(t *testing.T) {
	email := "tom@qq.com"
	value := url.Values{}
	value.Add("email", email)
	value.Add("password", "123456")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/login",
		bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), email)
	//assert.Equal(t, strings.Contains(w.Body.String(), email), true)
}