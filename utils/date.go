/*
通用工具包
*/
package utils

import (
	"database/sql"
	"time"
)

/*
@Description 获取 Mysql datetime 格式的当前日期
*/

func GetMysqlNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetSqlNullTime() sql.NullTime {
	return sql.NullTime{Time: time.Now(), Valid: true}
}
