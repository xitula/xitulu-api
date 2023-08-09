/*
@Description 待办清单路由
*/
package router

import (
	"xitulu/router/handler"

	"github.com/gin-gonic/gin"
)

/*
@Description 注册待办1ingrain路由
*/
func registerTodo(r *gin.Engine) {
	// 返回所有数据
	r.GET("/todo", handler.TodoGet)
	// 返回指定ID的数据
	r.GET("/todo/:id", handler.TodoGetOne)
	// 新增
	r.POST("/todo", handler.TodoAdd)
	// 更新
	r.PUT("/todo", handler.TodoUpdate)
	// 删除
	r.DELETE("/todo/:id", handler.TodoDelete)
}
