/*
代办清单数据库操作
*/
package model

import (
	"errors"
	"fmt"
	"log"

	t "xitulu/types"
	u "xitulu/util"
)

func SelectCount(table string, where string) (int64, error) {
	log.Println("where", where)
	var countSql string
	if where != "" {
		countSql = fmt.Sprintf("SELECT COUNT(*) as count FROM %s WHERE %s", table, where)
		log.Println("countSql", countSql)
		log.Println("countSql", countSql)
	} else {
		countSql = fmt.Sprintf("SELECT COUNT(*) as count FROM %s", table)
	}
	resultCount, errCount := dbQuery(countSql)
	if errCount != nil {
		log.Fatal(errCount)
		return 0, errCount
	}
	count := resultCount[0]["count"]
	return count.(int64), nil
}

/*
@Description 查询所有数据
*/
func SelectTodos() (interface{}, error) {
	const sql = "SELECT * FROM todos ORDER BY lastUpdateDate DESC"
	results, err := dbQuery(sql)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return results, nil
}

func SelectTodosPage(sql string, params ...interface{}) (interface{}, error) {
	results, err := dbQuery(sql, params)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	count, _ := SelectCount("todos", "")

	finalResult := map[string]any{
		"total": count,
		"list":  results,
	}

	return finalResult, nil
}

func SelectTodosPageByConditions(currentPage int64, pageSize int64, orderBy string, filterBy string) (interface{}, error) {
	var order string
	switch orderBy {
	case "create-desc":
		order = "createDate"
	case "update-desc":
		order = "lastUpdateDate"
	default:
		order = "createDate"
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

	var sql string
	var results interface{}
	var resErr error
	start := (currentPage - 1) * pageSize
	if done == -1 {
		sql = fmt.Sprintf("SELECT * FROM todos ORDER BY %s DESC LIMIT ?, ?", order)
		results, resErr = dbQuery(sql, start, pageSize)
	} else {
		sql = fmt.Sprintf("SELECT * FROM todos WHERE done = %d ORDER BY %s DESC LIMIT ?, ?", done, order)

		results, resErr = dbQuery(sql, start, pageSize)
	}

	if resErr != nil {
		log.Fatalln(resErr)
		return nil, resErr
	}

	var count int64
	if done == -1 {
		count, _ = SelectCount("todos", "")
	} else {
		count, _ = SelectCount("todos", fmt.Sprintf("done = %d", done))
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
	sql := "INSERT INTO todos (contant, description, done, createDate, lastUpdateDate) VALUES (?, ?, 0, ?, ?)"
	_, err := dbExec(sql, todo.Contant, todo.Description, createDate, createDate)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return err
}

/*
@Description 更新指定ID对应的条目
*/
func UpdateTodo(todo t.Todo) error {
	sql := "UPDATE todos SET contant = ?, description = ?, done = ?, lastUpdateDate = ? WHERE id = ?"
	lastUpdateDate := u.GetMysqlNow()
	_, err := dbExec(sql, todo.Contant, todo.Description, todo.Done, lastUpdateDate, todo.Id)
	if err != nil {
		return err
	}
	return nil
}

/*
@Description 删除指定ID的条目
*/
func DeleteTodo(id int64) error {
	sql := "DELETE FROM todos WHERE id = ?"
	rows, err := dbExec(sql, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if rows == 0 {
		return errors.New("id不存在")
	}
	return nil
}
