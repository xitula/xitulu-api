/*
代办清单数据库操作
*/
package model

import (
	"errors"
	"log"

	t "xitulu/types"
	u "xitulu/util"
)

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

func SelectTodosPage(currentPage int64, pageSize int64) (interface{}, error) {
	const sql = "SELECT * FROM todos ORDER BY lastUpdateDate DESC LIMIT ?, ?"
	start := (currentPage - 1) * pageSize
	results, err := dbQuery(sql, start, pageSize)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	const countSql = "SELECT COUNT(*) as count FROM todos"
	resultCount, errCount := dbQuery(countSql)
	if errCount != nil {
		log.Fatal(errCount)
		return nil, errCount
	}
	count := resultCount[0]["count"]

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
