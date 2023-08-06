package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"xitulu/models"
	"xitulu/utils"
)

var article models.Article

func init() {
	article = models.Article{}
}

// 新增文章
func ArticleAdd(c *gin.Context) {
	var article models.Article
	if errBind := c.ShouldBindJSON(&article); errBind != nil {
		response(c, errBind)
	}
	article.State = 1
	article.CreatedOn = utils.GetMysqlNow()

	id, err := article.Insert(&article)
	responseData(c, err, map[string]int{"id": id})
}

// 获取所有文章
func Articles(c *gin.Context) {
	if data, count, err := article.SelectAll(); err != nil {
		responseData(c, err, nil)
	} else {
		resp := map[string]interface{}{"count": count, "list": data}
		responseData(c, nil, resp)
	}
}

// 更新文章
func ArticleUpdate(c *gin.Context) {
	var article models.Article
	errBind := c.ShouldBindJSON(&article)
	if errBind != nil {
		response(c, errBind)
	} else {
		modifiedOn := utils.GetMysqlNow()
		article.ModifiedOn = &modifiedOn
		err := article.Update(&article)
		response(c, err)
	}
}

// 删除文章
func ArticleDelete(c *gin.Context) {
	sId := c.Param("id")
	id, _ := strconv.Atoi(sId)
	err := article.Delete(id)
	response(c, err)
}
