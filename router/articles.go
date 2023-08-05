package router

import (
	"xitulu/services"

	"github.com/gin-gonic/gin"
)

func addArticle(c *gin.Context) {
	var article services.Article
	if errBind := c.ShouldBindJSON(&article); errBind != nil {
		response(c, errBind)
	}
	errS := services.Add(&article)
	response(c, errS)
}

func getArticle(c *gin.Context)    {}
func getArticles(c *gin.Context)   {}
func updateArticle(c *gin.Context) {}
func deleteArticle(c *gin.Context) {}

func registerArticles(r *gin.Engine) {
	r.POST("/articles", addArticle)
	r.GET("/articles/:id", getArticle)
	r.GET("/articles", getArticles)
	r.PUT("/articles", updateArticle)
	r.DELETE("/articles", deleteArticle)
}
