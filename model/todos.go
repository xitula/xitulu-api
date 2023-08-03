/*
代办清单数据库操作
*/
package model

import (
	"errors"
	"log"

	t "xitulu/types"
	u "xitulu/util"

	"gorm.io/gorm/clause"
)

// 查询所有数据
func SelectTodos() interface{} {
	var todos []t.Todo
	orm.Table("todos").Where("status = 1").Find(&todos)
	return &todos
}

// 依据条件分页查询待办列表
func SelectTodosPageByConditions(currentPage int, pageSize int, orderBy string, filterBy string) (interface{}, error) {
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

	var results []t.Todo
	var count int64
	start := (currentPage - 1) * pageSize
	if done == -1 {
		orm.
			Table("todos").
			Where("status = 1").
			Count(&count).
			Order(clause.OrderByColumn{Column: clause.Column{Name: order}, Desc: true}).
			Offset(start).
			Limit(pageSize).
			Find(&results)
	} else {
		orm.
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
func SelectTodo(id int64) (interface{}, error) {
	sql := "SELECT * FROM todos WHERE id = ? AND status = 1"
	results, err := dbQuery(sql, id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &results, nil
}

// 插入新的待办条目
func InsertTodo(todo *t.Todo) error {
	createDate := u.GetMysqlNow()
	todo.Done = 0
	todo.CreateDate = createDate
	todo.LastUpdateDate = createDate
	result := orm.Table("todos").Create(&todo)

	if result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}
	return nil
}

// 更新指定ID对应的条目
func UpdateTodo(todo *t.Todo) error {
	lastUpdateDate := u.GetMysqlNow()
	result := orm.
		Table("todos").
		Where("id = ?", todo.Id).
		Updates(map[string]interface{}{"content": todo.Content, "description": todo.Description, "done": todo.Done, "last_update_date": lastUpdateDate})

	err := result.Error
	if err != nil {
		log.Fatalln("UpdateTodoError:", err)
		return err
	}
	return nil
}

// 删除指定ID的条目
func DeleteTodo(id int) error {
	lastUpdateDate := u.GetMysqlNow()
	result := orm.Table("todos").Where("id = ?", id).Updates(map[string]interface{}{"status": 0, "last_update_date": lastUpdateDate})

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
