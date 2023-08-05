package services

import (
	"xitulu/models"
	"xitulu/utils"
)

type Article struct {
	models.Article

	CurrentPage int
	PageSize    int
}

func Add(article *Article) error {
	article.CreatedOn = utils.GetMysqlNow()

	err := models.Article{}.InsertArticle(*article)

	return err
}
