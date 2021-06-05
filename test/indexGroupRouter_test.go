package initRouter

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 测试 indexGroupRouter.go
func TestIndexGroupRouter(t *testing.T) {
	w := httptest.NewRecorder()
	// 发送 Get 请求到 Any 路由
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "hello gin get method",
		"返回的HTML页面中应该包含 hello gin get method")

	w2 := httptest.NewRecorder()
	// 发送 Post 请求到 Any 路由
	req2, _ := http.NewRequest(http.MethodPost, "/", nil)
	router.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Contains(t, w2.Body.String(), "hello gin post method",
		"返回的HTML页面中应该包含 hello gin post method")

	w3 := httptest.NewRecorder()
	// 发送 Put 请求到 Any 路由
	req3, _ := http.NewRequest(http.MethodPut, "/", nil)
	router.ServeHTTP(w3, req3)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Contains(t, w3.Body.String(), "hello gin put method",
		"返回的HTML页面中应该包含 hello gin put method")

	w4 := httptest.NewRecorder()
	// 发送 Delete 请求到 Any 路由
	req4, _ := http.NewRequest(http.MethodDelete, "/", nil)
	router.ServeHTTP(w4, req4)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Contains(t, w4.Body.String(), "hello gin delete method",
		"返回的HTML页面中应该包含 hello gin delete method")
}
