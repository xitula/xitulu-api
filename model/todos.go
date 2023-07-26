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

/*
@Description 查询所有数据
*/
func SelectTodos() interface{} {
	var todos []t.Todo
	dbOrm.Table("todos").Find(&todos)
	return todos
}

func SelectTodosPageByConditions(currentPage int, pageSize int, orderBy string, filterBy string) (interface{}, error) {
	var order string
	switch orderBy {
	case "create-desc":
		order = "create_date"
	case "update-desc":
		order = "last_update_date"
	default:
		order = "create_date"
	}
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
		dbOrm.
			Table("todos").
			Count(&count).
			Order(clause.OrderByColumn{Column: clause.Column{Name: order}, Desc: true}).
			Offset(start).
			Limit(pageSize).
			Find(&results)
	} else {
		dbOrm.
			Table("todos").
			Where("done = ?", done).
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

	return finalResult, nil
}

/*
@Description 查询指定ID的条目
*/
func SelectTodo(id int64) (interface{}, error) {
	sql := "SELECT * FROM todos WHERE id = ?"
	results, err := dbQuery(sql, id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return results, nil
}

/*
@Description 插入新的待办条目
*/
func InsertTodo(todo t.Todo) error {
	createDate := u.GetMysqlNow()
	todo.Done = 0
	todo.CreateDate = createDate
	todo.LastUpdateDate = createDate
	result := dbOrm.Table("todos").Create(&todo)
	log.Println("todo", todo)
	if result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}
	return nil
}

/*
@Description 更新指定ID对应的条目
*/
func UpdateTodo(todo t.Todo) error {
	lastUpdateDate := u.GetMysqlNow()
	result := dbOrm.
		Table("todos").
		Where("id = ?", todo.Id).
		Updates(t.Todo{Content: todo.Content, Description: todo.Description, Done: todo.Done, LastUpdateDate: lastUpdateDate})

	err := result.Error
	if err != nil {
		log.Fatalln("UpdateTodoError:", err)
		return err
	}
	return nil
}

/*
@Description 删除指定ID的条目
*/
func DeleteTodo(id int) error {
	result := dbOrm.Table("todos").Delete(&t.Todo{}, id)

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
