package router

import (
	"github.com/gin-gonic/gin"
	"xitulu/models"
	"xitulu/router/handler"
)

// 新增文章
func addArticle(c *gin.Context) {
	var article models.Article
	if errBind := c.ShouldBindJSON(&article); errBind != nil {
		response(c, errBind)
	}
	errS := handler.ArticleAdd(&article)
	response(c, errS)
}

func getArticle(c *gin.Context) {}

// 获取所有文章
func getArticles(c *gin.Context) {
	if data, err := handler.ArticleGetAll(); err != nil {
		responseData(c, err, nil)
	} else {
		responseData(c, nil, data)
	}
}
func updateArticle(c *gin.Context) {}
func deleteArticle(c *gin.Context) {}

func registerArticles(r *gin.Engine) {
	r.POST("/articles", addArticle)
	r.GET("/articles/:id", getArticle)
	r.GET("/articles", getArticles)
	r.PUT("/articles", updateArticle)
	r.DELETE("/articles", deleteArticle)
}
