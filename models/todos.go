package models

import (
	"database/sql"
	"log"
)

// 待办模型
type Todo struct {
	Id             int           `json:"id" gorm:"column:id;primary"`                  // ID
	Uid            int           `json:"uid"`                                          // 用户ID
	Content        string        `json:"content"`                                      // 内容
	Description    *string       `json:"description,omitempty" gorm:"default:null"`    // 描述，可选
	Done           int           `json:"done,omitempty"`                               // 是否已完成
	DoneOn         *sql.NullTime `json:"doneOn,omitempty" gorm:"default:null"`         // 完成时间
	Status         int           `json:"status,omitempty" gorm:"default:1"`            // 条目状态
	CreateDate     sql.NullTime  `json:"createDate,omitempty"`                         // 创建日期
	LastUpdateDate *sql.NullTime `json:"lastUpdateDate,omitempty" gorm:"default:null"` // 最后更新日期，可选
}

func (t *Todo) Insert(todo *Todo) error {
	result := db.Table("todos").Create(&todo)

	if result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}
	return nil
}

func (t *Todo) Update(todo *Todo) error {
	result := db.
		Table("todos").
		Model(&todo).
		Select("Done").
		Updates(&todo)

	err := result.Error
	if err != nil {
		log.Fatalln("UpdateTodoError:", err)
		return err
	}
	return nil
}
