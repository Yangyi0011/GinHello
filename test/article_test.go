package initRouter

import (
	"GinHello/model"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 测试文章数据插入
func TestArticleInsert(t *testing.T) {
	article := model.Article{
		Type:    "c++",
		Content: "hello c++",
	}
	marshal, err := json.Marshal(article)
	if err != nil {
		log.Panicln("json转化错误：", err)
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/article", bytes.NewBuffer(marshal))
	req.Header.Add("context-type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.NotEqual(t, "{\"id\":-1}", w.Body.String())
}
