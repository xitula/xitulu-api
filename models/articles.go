package models

// 文章模型
type Article struct {
	// 文章ID
	Id int `json:"id" gorm:"primary"`
	// 创建人ID
	Uid int `json:"uid"`
	// 标题
	Title string `json:"title"`
	// 描述
	Description *string `json:"description"`
	// 内容
	Content string `json:"content"`
	// 状态 0=删除 1=正常
	State int `json:"state"`
	// 创建时间
	CreatedOn string `json:"createdOn,omitempty"`
	// 修改时间
	ModifiedOn *string `json:"modifiedOn,omitempty"`
}

// 插入文章
func InsertArticle(data *Article) error {
	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

// 查询所有文章
func SelectAll() (*[]Article, int64, error) {
	var articles *[]Article
	var count int64
	if err := db.Table("articles").Find(&articles).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return articles, count, nil
}
