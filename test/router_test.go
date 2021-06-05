package initRouter

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

/*
	此文件已不再使用，是之前学习时的记录
*/
func TestIndexGetRouter(t *testing.T) {
	// 获取响应对象
	w := httptest.NewRecorder()
	// 设置请求方法、请求拦截路径、请求包体，以获取请求对象
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	// 给路由绑定响应对象和请求对象
	router.ServeHTTP(w, req)
	// 断言状态码
	assert.Equal(t, http.StatusOK, w.Code)
	//断言响应数据
	assert.Equal(t, "hello gin get method", w.Body.String())
}

// 测试 Post 请求
func TestIndexPostRouter(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	router.ServeHTTP(w, req)
	// 断言状态码
	assert.Equal(t, http.StatusOK, w.Code)
	//断言响应数据
	assert.Equal(t, "hello gin post method", w.Body.String())
}

// 测试 Put 请求
func TestIndexPutRouter(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/", nil)
	router.ServeHTTP(w, req)
	// 断言状态码
	assert.Equal(t, http.StatusOK, w.Code)
	//断言响应数据
	assert.Equal(t, "hello gin put method", w.Body.String())
}

// 测试 Delete 请求
func TestIndexDeleteRouter(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/", nil)
	router.ServeHTTP(w, req)
	// 断言状态码
	assert.Equal(t, http.StatusOK, w.Code)
	//断言响应数据
	assert.Equal(t, "hello gin delete method", w.Body.String())
}

// 测试 Patch 请求
func TestIndexPatchRouter(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPatch, "/", nil)
	router.ServeHTTP(w, req)
	// 断言状态码
	assert.Equal(t, http.StatusOK, w.Code)
	//断言响应数据
	assert.Equal(t, "hello gin patch method", w.Body.String())
}

// 测试 Head 请求
func TestIndexHeadRouter(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodHead, "/", nil)
	router.ServeHTTP(w, req)
	// 断言状态码
	assert.Equal(t, http.StatusOK, w.Code)
	//断言响应数据
	assert.Equal(t, "hello gin head method", w.Body.String())
}

// 测试 Options 请求
func TestIndexOptionsRouter(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodOptions, "/", nil)
	router.ServeHTTP(w, req)
	// 断言状态码
	assert.Equal(t, http.StatusOK, w.Code)
	//断言响应数据
	assert.Equal(t, "hello gin options method", w.Body.String())
}

// 测试路径参数
func TestUserSave(t *testing.T) {
	username := "tom"
	w := httptest.NewRecorder()
	// Restful 风格请求方式
	// 注意，虽然路由指定的路径是 /user/:name，但这里请求时不需要 :
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户【"+username+"】已保存", w.Body.String())
}

// 测试通过 context.Query() 获取路径参数
func TestUserSaveByQuery(t *testing.T) {
	username := "tom"
	age := 18
	w := httptest.NewRecorder()
	// 传统的 Get 方法请求方式
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+strconv.Itoa(age), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户："+username+"，年龄："+strconv.Itoa(age)+"已经保存", w.Body.String())
}

// 测试路由分组
func TestRouterGroup(t *testing.T) {
	// w 不能复用
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	// 断言状态码
	assert.Equal(t, http.StatusOK, w.Code)
	//断言响应数据
	assert.Equal(t, "hello gin get method", w.Body.String())

	username := "jerry"
	// 复用 w 会导致断言不通过
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, "用户【"+username+"】已保存", w2.Body.String())

	age := 20
	// 复用 w 会导致断言不通过
	w3 := httptest.NewRecorder()
	req3, _ := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+strconv.Itoa(age), nil)
	router.ServeHTTP(w3, req3)
	assert.Equal(t, http.StatusOK, w3.Code)
	assert.Equal(t, "用户："+username+"，年龄："+strconv.Itoa(age)+"已经保存", w3.Body.String())
}

// 测试 Any 请求方法
func TestAnyRouter(t *testing.T) {
	w := httptest.NewRecorder()
	// 发送 Get 请求到 Any 路由
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin get method", w.Body.String())

	w2 := httptest.NewRecorder()
	// 发送 Post 请求到 Any 路由
	req2, _ := http.NewRequest(http.MethodPost, "/", nil)
	router.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, "hello gin post method", w2.Body.String())

	w3 := httptest.NewRecorder()
	// 发送 Put 请求到 Any 路由
	req3, _ := http.NewRequest(http.MethodPut, "/", nil)
	router.ServeHTTP(w3, req3)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, "hello gin put method", w3.Body.String())

	w4 := httptest.NewRecorder()
	// 发送 Delete 请求到 Any 路由
	req4, _ := http.NewRequest(http.MethodDelete, "/", nil)
	router.ServeHTTP(w4, req4)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Equal(t, "hello gin delete method", w4.Body.String())
}

// 测试 index.tmpl 页面路由
func TestIndexHTML(t *testing.T) {
	// 这里如果没有设置 ginMode=TestMode，测试时会导致找不到 index.tmpl 文件
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "hello gin get method","返回的HTML页面中应该包含 hello gin get method")
}