package models

import (
	"errors"
	"log"
	"time"
	t "xitulu/types"
	u "xitulu/utils"
)

// 随感
type Causerie struct {
	Id             int        `json:"id,omitempty" gorm:"column:id;primary"`
	Uid            int        `json:"uid,omitempty" gorm:"column:uid"`
	Content        string     `json:"content,omitempty" gorm:"column:content"`
	Status         int        `json:"status,omitempty" gorm:"column:status"`
	CreateDate     time.Time  `json:"createDate,omitempty" gorm:"column:create_date;default:null"`
	LastUpdateDate *time.Time `json:"lastUpdateDate,omitempty" gorm:"column:last_update_date;default:null"`
}

// SelectAll 获取全部随感
func (c *Causerie) SelectAll() ([]Causerie, error) {
	var data []Causerie
	result := db.Table("causeries").Where("status = 1").Order("create_date DESC").Find(&data)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return data, nil
	}
}

// SelectAllPage 分页获取全部随感
func (c *Causerie) SelectAllPage(page *t.Pagination) ([]Causerie, int64, error) {
	var data []Causerie
	var count int64
	offset := (page.CurrentPage - 1) * page.PageSize
	result := db.
		Table("causeries").
		Where("status = 1").
		Count(&count).
		Offset(offset).
		Limit(page.PageSize).
		Find(&data)
	if result.Error != nil {
		return nil, 0, result.Error
	} else {
		return data, count, nil
	}
}

// Insert 新增随感
func (c *Causerie) Insert(data *Causerie) error {
	result := db.Table("causeries").Create(&data)
	if err := result.Error; err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

// Update 依据ID更新随感数据
func (c *Causerie) Update(id int, content string) error {
	lastUpdateDate := u.GetMysqlNow()
	result := db.Table("causeries").Where("id = ?", id).Updates(t.Causerie{Id: id, Content: content, LastUpdateDate: lastUpdateDate})
	errUpd := result.Error
	if errUpd != nil {
		log.Fatalln("UpdateCauserieError:", errUpd)
	}
	if result.RowsAffected == 0 {
		return errors.New("id错误")
	}

	return nil
}

// Delete 依据ID删除随感
func (c *Causerie) Delete(id int) error {
	lastUpdateDate := u.GetMysqlNow()
	result := db.Table("causeries").Where("id = ?", id).
		Updates(map[string]interface{}{"id": id, "status": 0, "last_update_date": lastUpdateDate})
	errUpd := result.Error
	if errUpd != nil {
		log.Fatalln("UpdateCauserieError:", errUpd)
	}
	if result.RowsAffected == 0 {
		return errors.New("id错误")
	}

	return nil
}
