package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"time"
	"xitulu/models"
)

var todo models.Todo

func init() {
	todo = models.Todo{}
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
