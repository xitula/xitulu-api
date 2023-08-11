package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"xitulu/models"
	"xitulu/types"
)

var modelTodo models.Todo

func init() {
	modelTodo = models.Todo{}
}

func TodoGet(ctx *gin.Context) {
	var params types.GetAllParam
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response(ctx, err)
		return
	}

	if err := validate.Struct(&params); err != nil {
		response(ctx, err)
		return
	}

	if params.CurrentPage != 0 && params.PageSize != 0 {
		data, errPage := modelTodo.SelectByConditions(&params)
		if errPage != nil {
			response(ctx, errPage)
		} else {
			responseData(ctx, errPage, data)
		}
	} else {
		data := modelTodo.SelectAll()
		responseData(ctx, nil, data)
	}
}

func TodoGetOne(ctx *gin.Context) {
	sId := ctx.Param("id")
	id, errId := strconv.Atoi(sId)
	if errId != nil {
		response(ctx, errId)
		return
	}
	if errs := validate.Var(id, "required,gt=0"); errs != nil {
		response(ctx, errs)
		return
	}

	data, err := modelTodo.SelectTodo(id)
	responseData(ctx, err, data)
}

func TodoAdd(ctx *gin.Context) {
	var data models.Todo
	errBind := ctx.BindJSON(&data)
	if errBind != nil {
		response(ctx, errBind)
		return
	}
	if errs := validate.Struct(&data); errs != nil {
		response(ctx, errs)
		return
	}
	data.Done = 0
	data.CreateDate = &sql.NullTime{Time: time.Now(), Valid: true}
	data.Status = 1

	err := modelTodo.Insert(&data)
	response(ctx, err)
}

func TodoUpdate(ctx *gin.Context) {
	var data models.Todo
	errBind := ctx.BindJSON(&data)
	if errBind != nil {
		response(ctx, errBind)
		return
	}
	if errs := validate.Struct(&data); errs != nil {
		response(ctx, errs)
		return
	}
	now := time.Now()
	data.LastUpdateDate = &sql.NullTime{Time: now, Valid: true}
	err := modelTodo.Update(&data)
	response(ctx, err)
}

func TodoDelete(ctx *gin.Context) {
	sId := ctx.Param("id")
	id, errId := strconv.Atoi(sId)
	if errId != nil {
		response(ctx, errId)
	}
	if errs := validate.Var(id, "gt=0"); errs != nil {
		response(ctx, errs)
		return
	}
	errExec := modelTodo.Delete(id)
	response(ctx, errExec)
}
