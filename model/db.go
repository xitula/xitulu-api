/*
@Description 数据库相关操作函数
*/
package model

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

/*
@Description 配置并初始化数据库
*/
func SetupDb() {
	var err error
	db, err = sql.Open("mysql", "root:yl123456@tcp(localhost:3306)/xitulu")

	if err != nil {
		panic(err.Error)
	}
}

/*
@Description 执行查询并返回数据
*/
func dbQuery(sql string, params ...interface{}) (interface{}, error) {
	trans, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	stmt, err := trans.Prepare(sql)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	rows, err := stmt.Query(params...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
				if nx, ok := strconv.ParseFloat(string(b), 64); ok == nil {
					v = nx
				}
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	commitErr := trans.Commit()
	if commitErr != nil {
		return nil, commitErr
	}
	return tableData, nil
}

/*
@Description 执行SQL并返回结果
*/
func dbExec(sql string, params ...interface{}) (int64, error) {
	trans, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}
	stmt, err := trans.Prepare(sql)
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}
	res, err := stmt.Exec(params...)
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}
	affectedCount, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	trans.Commit()
	return affectedCount, nil
}

/*
@Description 关闭数据库连接
*/
func CloseDb() {
	defer db.Close()
}
