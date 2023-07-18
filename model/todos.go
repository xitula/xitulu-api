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
	results, err := dbQuery("SELECT * FROM todos ORDER BY createDate DESC")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return results, nil
}

/*
@Description 查询指定ID的条目
*/
func SelectTodo(params t.ReqParams) (interface{}, error) {
	sql := makeSql(`SELECT * FROM todos WHERE id = ${id}`, params)
	results, err := dbQuery(sql)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return results, nil
}

/*
@Description 插入新的待办条目
*/
func InsertTodo(params t.ReqParams) error {
	createDate := u.GetMysqlNow()
	params["createDate"] = createDate
	sql := makeSql(`INSERT INTO todos (contant, description, createDate) VALUES ("${contant}", "${description}", "${createDate}")`, params)
	_, err := dbExec(sql)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return err
}

/*
@Description 更新指定ID对应的条目
*/
func UpdateTodo(params t.ReqParams) error {
	lastUpdateDate := u.GetMysqlNow()
	params["lastUpdateDate"] = lastUpdateDate
	sql := makeSql(`UPDATE todos SET contant = "${contant}", description = "${description}", lastUpdateDate = "${lastUpdateDate}" WHERE id = ${id}`, params)
	rows, err := dbExec(sql)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if rows == 0 {
		return errors.New("id不存在")
	}
	return nil
}

/*
@Description 删除指定ID的条目
*/
func DeleteTodo(params t.ReqParams) error {
	sql := makeSql(`DELETE FROM todos WHERE id = ${id}`, params)
	rows, err := dbExec(sql)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if rows == 0 {
		return errors.New("id不存在")
	}
	return nil
}
