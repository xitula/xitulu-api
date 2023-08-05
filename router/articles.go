package router

import (
	"github.com/gin-gonic/gin"
	"xitulu/router/handler"
)

func registerArticles(r *gin.Engine) {
	r.POST("/articles", handler.ArticleAdd)
	//r.GET("/articles/:id", getArticle)
	r.GET("/articles", handler.Articles)
	r.PUT("/articles", handler.ArticleUpdate)
	r.DELETE("/articles/:id", handler.ArticleDelete)
}
