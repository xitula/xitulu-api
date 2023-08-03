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
	// 查询所有用户
	r.GET("/users", func(ctx *gin.Context) {
		data, err := model.SelectUsersAll()
		responseData(ctx, err, data)
	})

	// 新增用户
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

	// 更新用户数据
	r.PUT("/users", func(ctx *gin.Context) {
		var user t.UserModel
		errBind := ctx.ShouldBindJSON(&user)
		if errBind != nil {
			log.Fatalln("errBind", errBind)
			response(ctx, errBind)
			return
		}
		err := model.UpdateUser(&user)
		response(ctx, err)
	})

	// 删除用户
	r.DELETE("/users/:id", func(ctx *gin.Context) {
		sId := ctx.Param("id")
		id, _ := strconv.Atoi(sId)
		err := model.DeleteUser(id)
		response(ctx, err)
	})

	// 用户登录
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
			token, _ := model.UpdateUserUuid(dbUser.Id, false)
			responseData(ctx, err, map[string]interface{}{"id": dbUser.Id, "username": dbUser.Username, "nickname": dbUser.Nickname, "email": dbUser.Email, "token": token})
		}
	})

	// 用户登出
	r.GET("/users/logout", func(ctx *gin.Context) {
		sId := ctx.Query("id")
		id, _ := strconv.Atoi(sId)
		_, err := model.UpdateUserUuid(id, true)
		response(ctx, err)
	})
}
