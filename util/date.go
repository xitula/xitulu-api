/*
通用工具包
*/
package util

import "time"

/*
@Description 获取 Mysql datetime 格式的当前日期
*/
func GetMysqlNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
