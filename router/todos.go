/*
@Description 待办清单路由
*/
package router

import (
	"fmt"
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
		data, err := model.SelectTodos()
		responseData(ctx, err, data)
	})

	// 返回指定ID的数据
	r.GET("/todo/:id", func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		params := t.ReqParams{"id": id}
		data, err := model.SelectTodo(params)
		responseData(ctx, err, data)
	})

	// 新增
	r.POST("/todo", func(ctx *gin.Context) {
		contant := ctx.PostForm("contant")
		description := ctx.PostForm("description")
		params := map[string]string{"contant": contant, "description": description}
		fmt.Println(params)
		err := model.InsertTodo(params)
		response(ctx, err)
	})

	// 更新
	r.PUT("/todo/:id", func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		contant := ctx.PostForm("contant")
		description := ctx.PostForm("description")
		params := map[string]string{"id": id, "contant": contant, "description": description}
		fmt.Println(params)
		err := model.UpdateTodo(params)
		response(ctx, err)
	})

	// 删除
	r.DELETE("/todo/:id", func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		params := t.ReqParams{"id": id}
		err := model.DeleteTodo(params)
		response(ctx, err)
	})
}
