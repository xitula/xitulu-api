/*
@Description 数据库相关操作函数
*/
package model

import (
	"database/sql"
	"os"

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

func makeSql(sql string, params map[string]string) string {
	return os.Expand(sql, func(key string) string { return params[key] })
}

/*
@Description 执行查询并返回数据
*/
func dbQuery(sqlString string) (interface{}, error) {
	rows, err := db.Query(sqlString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
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
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	return tableData, nil
}

/*
@Description 执行SQL并返回结果
*/
func dbExec(sql string) (int64, error) {
	result, err := db.Exec(sql)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}

/*
@Description 关闭数据库连接
*/
func CloseDb() {
	defer db.Close()
}
