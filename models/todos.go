package models

import (
	"database/sql"
	"errors"
	"gorm.io/gorm/clause"
	"log"
	u "xitulu/utils"
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

// 依据条件分页查询待办列表
func (t *Todo) SelectByConditions(currentPage int, pageSize int, orderBy string, filterBy string) (interface{}, error) {
	// 排序条件
	var order string
	switch orderBy {
	case "create-desc":
		order = "create_date"
	case "update-desc":
		order = "last_update_date"
	default:
		order = "create_date"
	}
	// 是否已完成
	var done int8
	switch filterBy {
	case "tobe":
		done = 0
	case "done":
		done = 1
	default:
		done = -1
	}

	var results []Todo
	var count int64
	start := (currentPage - 1) * pageSize
	if done == -1 {
		db.
			Table("todos").
			Where("status = 1").
			Count(&count).
			Order(clause.OrderByColumn{Column: clause.Column{Name: order}, Desc: true}).
			Offset(start).
			Limit(pageSize).
			Find(&results)
	} else {
		db.
			Table("todos").
			Where("done = ? AND status = 1", done).
			Count(&count).
			Order(clause.OrderByColumn{Column: clause.Column{Name: order}, Desc: true}).
			Offset(int(start)).
			Limit(int(pageSize)).
			Find(&results)
	}

	finalResult := map[string]any{
		"total": count,
		"list":  results,
	}

	return &finalResult, nil
}

// 查询指定ID的条目
func (t *Todo) SelectTodo(id int) (*Todo, error) {
	var data Todo
	if err := db.Table("todo").Where("id = ? AND status = 1", id).Take(&data).Error; err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		return &data, nil
	}
}

// 查询所有数据
func (t *Todo) SelectAll() interface{} {
	var todos []Todo
	db.Table("todos").Where("status = 1").Find(&todos)
	return &todos
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

// 删除指定ID的条目
func (t *Todo) Delete(id int) error {
	lastUpdateDate := u.GetSqlNullTime()
	result := db.
		Table("todos").
		Where("id = ?", id).
		Updates(Todo{Status: 0, LastUpdateDate: &lastUpdateDate})

	err := result.Error
	if err != nil {
		log.Fatalln(err)
		return err
	}
	if result.RowsAffected == 0 {
		return errors.New("id不存在")
	}
	return nil
}
