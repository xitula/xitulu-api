package handler

import (
	"xitulu/models"
	"xitulu/utils"
)

// 新增文章
func ArticleAdd(article *models.Article) error {
	article.State = 1
	article.CreatedOn = utils.GetMysqlNow()

	err := models.InsertArticle(article)

	return err
}

// 获取所有文章
func ArticleGetAll() (*map[string]interface{}, error) {
	data, count, err := models.SelectAll()
	return &map[string]interface{}{"count": count, "list": data}, err
}
