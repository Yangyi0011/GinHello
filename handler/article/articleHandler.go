package article

import (
	"GinHello/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Insert godoc
// @Summary 添加文章
// @Description 提交新的文章内容
// @Tags 文章管理
// @version 1.0
// @Accept json
// @Produce  json
// @Param article body model.Article true "文章"
// @Success 200 object model.Result 成功后返回值
// @Failure 409 object model.Result 参数格式有误
// @Failure 500 object model.Result 添加失败
// @Router /article [post]
func Insert(context *gin.Context) {
	var article model.Article
	if e := context.ShouldBindJSON(&article); e != nil {
		context.JSON(http.StatusConflict, gin.H{
			"result": model.Result{
				Code:    http.StatusConflict,
				Message: "参数格式有误",
				Data:    e.Error(),
			},
		})
		return
	}
	id, e := article.Insert()
	if e != nil {
		context.JSON(http.StatusConflict, gin.H{
			"result": model.Result{
				Code:    http.StatusInternalServerError,
				Message: "添加失败",
				Data:    e.Error(),
			},
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"result": model.Result{
			Code:    http.StatusOK,
			Message: "添加成功",
			Data:    id,
		},
	})
}

// Update godoc
// @Summary 更新文章
// @Description 更新文章内容
// @Tags 文章管理
// @version 1.0
// @Accept json
// @Produce  json
// @Param article body model.Article true "文章"
// @Success 200 object model.Result 成功后返回值
// @Failure 409 object model.Result 参数格式有误
// @Failure 500 object model.Result 更新失败
// @Router /article [put]
func Update(context *gin.Context) {
	var article model.Article
	if e := context.ShouldBindJSON(&article); e != nil {
		context.JSON(http.StatusConflict, gin.H{
			"result": model.Result{
				Code:    http.StatusConflict,
				Message: "参数格式有误",
				Data:    e.Error(),
			},
		})
		return
	}
	e := article.Update()
	if e != nil {
		context.JSON(http.StatusConflict, gin.H{
			"result": model.Result{
				Code:    http.StatusInternalServerError,
				Message: "更新失败",
				Data:    e.Error(),
			},
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"result": model.Result{
			Code:    http.StatusOK,
			Message: "更新成功",
			Data:    article,
		},
	})
}

// GetOne godoc
// @Summary 查询文章
// @Description 通过 id 查询指定文章
// @Tags 文章管理
// @version 1.0
// @Accept json
// @Produce  json
// @Param id path int true "id"
// @Success 200 object model.Result 成功后返回值
// @Failure 400 object model.Result 查询失败
// @Router /article/{id} [get]
func GetOne(context *gin.Context) {
	id := context.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"result": model.Result{
				Code:    http.StatusBadRequest,
				Message: "参数格式有误",
				Data:    e.Error(),
			},
		})
		return
	}
	article := model.Article{
		Id: i,
	}
	art, e := article.FindById()
	if e != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"result": model.Result{
				Code:    http.StatusBadRequest,
				Message: "查询失败",
				Data:    e.Error(),
			},
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"result": model.Result{
			Code:    http.StatusOK,
			Message: "查询成功",
			Data:    art,
		},
	})
}

// GetAll godoc
// @Summary 查询所有文章
// @Description 查询所有文章
// @Tags 文章管理
// @version 1.0
// @Accept json
// @Produce  json
// @Success 200 object model.Result 成功后返回值
// @Router /article [get]
func GetAll(context *gin.Context) {
	article := model.Article{}
	articles, e := article.FindAll()
	if e != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"result": model.Result{
				Code:    http.StatusBadRequest,
				Message: "查询失败",
				Data:    e.Error(),
			},
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"result": model.Result{
			Code:    http.StatusOK,
			Message: "查询成功",
			Data:    articles,
		},
	})
}

// DeleteOne godoc
// @Summary 删除文章
// @Description 通过 id 删除指定文章
// @Tags 文章管理
// @version 1.0
// @Accept json
// @Produce  json
// @Param id path int true "id"
// @Success 200 object model.Result 成功后返回值
// @Failure 400 object model.Result 删除失败
// @Router /article/{id} [delete]
func DeleteOne(context *gin.Context) {
	id := context.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"result": model.Result{
				Code:    http.StatusBadRequest,
				Message: "参数格式有误",
				Data:    e.Error(),
			},
		})
		return
	}
	article := model.Article{Id: i}
	e = article.DeleteOne()
	if e != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"result": model.Result{
				Code:    http.StatusBadRequest,
				Message: "删除失败",
				Data:    e.Error(),
			},
		})
		return
	}
	context.JSON(http.StatusBadRequest, gin.H{
		"result": model.Result{
			Code:    http.StatusOK,
			Message: "删除成功",
			Data:    id,
		},
	})
}