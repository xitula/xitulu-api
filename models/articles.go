package models

type Article struct {
	Id          int    `json:"id"`
	Uid         int    `json:"uid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	State       int    `json:"state"`
	CreatedOn   string `json:"createdOn"`
	ModifiedOn  string `json:"modifiedOn"`
}

func (a *Article) InsertArticle(data *map[string]interface{}) error {
	//article := Article{
	//	Uid:         data["uid"].(int),
	//	Title:       data["title"].(string),
	//	Description: data["description"].(string),
	//	Content:     data["content"].(string),
	//	State:       data["state"].(int),
	//	CreatedOn:   data["createOn"].(string),
	//}
	if err := db.Create(&article).Error; err != nil {
		return err
	}

	return nil
}
