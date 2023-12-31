package models

import (
	"errors"
	"xitulu/utils"
)

// 文章模型
type Article struct {
	// 文章ID
	Id int `json:"id,omitempty" gorm:"primary"`
	// 创建人ID
	Uid int `json:"uid" validate:"required,gt=0"`
	// 标题
	Title string `json:"title" validate:"required,min=1,max=255"`
	// 描述
	Description *string `json:"description,omitempty" validate:"min=1,max=255"`
	// 内容
	Content string `json:"content" validate:"required,min=1,max=65535"`
	// 状态 0=删除 1=正常
	State int `json:"state,omitempty"`
	// 创建时间
	CreatedOn string `json:"createdOn,omitempty"`
	// 修改时间
	ModifiedOn *string `json:"modifiedOn,omitempty"`
}

// 插入文章
func (a *Article) Insert(data *Article) (int, error) {
	if err := db.Create(data).Error; err != nil {
		return 0, err
	}

	return data.Id, nil
}

// 查询所有文章
func (a *Article) SelectAll() (*[]Article, int64, error) {
	var articles *[]Article
	var count int64
	if err := db.Table("articles").Where("state = 1").Order("created_on DESC").Find(&articles).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return articles, count, nil
}

// 更新文章
func (a *Article) Update(article *Article) error {
	result := db.
		Table("articles").
		Model(&article).
		Where("id = ?", article.Id).
		Select("Title", "Description", "Content", "ModifiedOn").
		Updates(&article)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("ID不存在")
	}
	return nil
}

// 删除文章
func (a *Article) Delete(id int) error {
	modifiedOn := utils.GetMysqlNow()
	err := db.Table("articles").Where("id = ?", id).Updates(map[string]interface{}{"state": 0, "modified_on": modifiedOn}).Error
	return err
}
