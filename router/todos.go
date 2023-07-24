/*
@Description 待办清单路由
*/
package router

import (
	"strconv"
	"xitulu/model"

	t "xitulu/types"

	"github.com/gin-gonic/gin"
)

/*
@Description 注册待办1ingrain路由
*/
func registerTodo(r *gin.Engine) {
	// 返回所有数据
	r.GET("/todo", func(ctx *gin.Context) {
		sCurrentPage := ctx.Query("currentPage")
		sPageSize := ctx.Query("pageSize")
		orderBy := ctx.Query("orderBy")
		filterBy := ctx.Query("filterBy")

		if sCurrentPage != "" && sPageSize != "" {
			currentPage, errCrr := strconv.ParseInt(sCurrentPage, 10, 64)
			if errCrr != nil {
				response(ctx, errCrr)
				return
			}
			pageSize, errSize := strconv.ParseInt(sPageSize, 10, 64)
			if errSize != nil {
				response(ctx, errSize)
				return
			}
			data, errPage := model.SelectTodosPageByConditions(currentPage, pageSize, orderBy, filterBy)
			if errPage != nil {
				response(ctx, errPage)
			} else {
				responseData(ctx, errPage, data)
			}
		} else {
			data, err := model.SelectTodos()
			responseData(ctx, err, data)
		}
	})

	// 返回指定ID的数据
	r.GET("/todo/:id", func(ctx *gin.Context) {
		sId := ctx.Params.ByName("id")
		id, errId := strconv.ParseInt(sId, 10, 64)
		if errId != nil {
			response(ctx, errId)
		}
		data, err := model.SelectTodo(id)
		responseData(ctx, err, data)
	})

	// 新增
	r.POST("/todo", func(ctx *gin.Context) {
		todo := t.Todo{}
		errBind := ctx.BindJSON(&todo)
		if errBind != nil {
			response(ctx, errBind)
			return
		}
		err := model.InsertTodo(todo)
		response(ctx, err)
	})

	// 更新
	r.PUT("/todo", func(ctx *gin.Context) {
		todo := t.Todo{}
		errBind := ctx.BindJSON(&todo)
		if errBind != nil {
			response(ctx, errBind)
			return
		}
		err := model.UpdateTodo(todo)
		response(ctx, err)
	})

	// 删除
	r.DELETE("/todo/:id", func(ctx *gin.Context) {
		sId := ctx.Params.ByName("id")
		id, errId := strconv.ParseInt(sId, 10, 64)
		if errId != nil {
			response(ctx, errId)
		}
		errExec := model.DeleteTodo(id)
		response(ctx, errExec)
	})
}
