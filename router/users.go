package router

import (
	"errors"
	"log"
	"strconv"
	"xitulu/model"
	t "xitulu/types"

	"github.com/gin-gonic/gin"
)

func registerUsers(r *gin.Engine) {
	r.GET("/users", func(ctx *gin.Context) {
		data, err := model.SelectUsersAll()
		responseData(ctx, err, data)
	})
	r.POST("/users", func(ctx *gin.Context) {
		var user t.UserAdd
		errBind := ctx.ShouldBindJSON(&user)
		if errBind != nil {
			log.Fatalln("errBind", errBind)
			response(ctx, errBind)
			return
		}
		data, err := model.InsertUser(&user)
		responseData(ctx, err, data)
	})
	r.PUT("/users", func(ctx *gin.Context) {
		data, err := model.SelectUsersAll()
		responseData(ctx, err, data)
	})
	r.DELETE("/users", func(ctx *gin.Context) {
		data, err := model.SelectUsersAll()
		responseData(ctx, err, data)
	})

	r.POST("/users/login", func(ctx *gin.Context) {
		var user t.UserLogin
		errBind := ctx.ShouldBindJSON(&user)
		if errBind != nil {
			log.Fatalln("errBind:", errBind)
			response(ctx, errBind)
			return
		}

		dbUser, err := model.SelectUserFirst(&user)

		if user.Password != dbUser.Password {
			response(ctx, errors.New("用户名或密码错误"))
			model.UpdateUserUuid(dbUser.Id, true)
		} else {
			str, _ := model.UpdateUserUuid(dbUser.Id, false)
			responseData(ctx, err, map[string]interface{}{"id": dbUser.Id, "username": dbUser.Username, "nickname": dbUser.Nickname, "email": dbUser.Email, "token": str})
		}
	})

	r.GET("/users/logout", func(ctx *gin.Context) {
		sId := ctx.Query("id")
		id, _ := strconv.Atoi(sId)
		_, err := model.UpdateUserUuid(id, true)
		response(ctx, err)
	})
}
