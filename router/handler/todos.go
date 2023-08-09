package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"xitulu/models"
)

var todo models.Todo

func init() {
	todo = models.Todo{}
}

func TodoGet(ctx *gin.Context) {
	sCurrentPage := ctx.Query("currentPage")
	sPageSize := ctx.Query("pageSize")
	orderBy := ctx.Query("orderBy")
	filterBy := ctx.Query("filterBy")

	if sCurrentPage != "" && sPageSize != "" {
		currentPage, errCrr := strconv.Atoi(sCurrentPage)
		if errCrr != nil {
			response(ctx, errCrr)
			return
		}
		pageSize, errSize := strconv.Atoi(sPageSize)
		if errSize != nil {
			response(ctx, errSize)
			return
		}
		data, errPage := todo.SelectByConditions(currentPage, pageSize, orderBy, filterBy)
		if errPage != nil {
			response(ctx, errPage)
		} else {
			responseData(ctx, errPage, data)
		}
	} else {
		data := todo.SelectAll()
		responseData(ctx, nil, data)
	}
}

func TodoGetOne(ctx *gin.Context) {
	sId := ctx.Param("id")
	id, errId := strconv.Atoi(sId)
	if errId != nil {
		response(ctx, errId)
	}
	data, err := todo.SelectTodo(id)
	responseData(ctx, err, data)
}

func TodoAdd(ctx *gin.Context) {
	var data models.Todo
	errBind := ctx.BindJSON(&data)
	if errBind != nil {
		response(ctx, errBind)
		return
	}
	data.Done = 0
	data.CreateDate = sql.NullTime{Time: time.Now(), Valid: true}
	data.Status = 1

	err := todo.Insert(&data)
	response(ctx, err)
}

func TodoUpdate(ctx *gin.Context) {
	var data models.Todo
	errBind := ctx.BindJSON(&data)
	if errBind != nil {
		response(ctx, errBind)
		return
	}
	now := time.Now()
	data.LastUpdateDate = &sql.NullTime{Time: now, Valid: true}
	err := todo.Update(&data)
	response(ctx, err)
}

func TodoDelete(ctx *gin.Context) {
	sId := ctx.Param("id")
	id, errId := strconv.Atoi(sId)
	if errId != nil {
		response(ctx, errId)
	}
	errExec := todo.Delete(id)
	response(ctx, errExec)
}
