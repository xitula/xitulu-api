package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"xitulu/models"
)

var modelUser models.User

func init() {
	modelUser = models.User{}
}

// UserGetAll 查询所有用户
func UserGetAll(ctx *gin.Context) {
	data, err := modelUser.SelectAll()
	responseData(ctx, err, data)
}

// UserAdd 新增用户
func UserAdd(ctx *gin.Context) {
	var data models.User
	if err := ctx.ShouldBindJSON(&data); err != nil {
		response(ctx, err)
		log.Fatalln("err", err)
	} else {
		result, err := modelUser.Insert(&data)
		responseData(ctx, err, result)
	}
}

// UserUpdate 更新用户数据
func UserUpdate(ctx *gin.Context) {
	var data models.User
	if errBind := ctx.ShouldBindJSON(&data); errBind != nil {
		response(ctx, errBind)
		log.Fatalln("errBind", errBind)
	} else {
		err := modelUser.Update(&data)
		response(ctx, err)
	}
}

// UserDelete 删除用户
func UserDelete(ctx *gin.Context) {
	sId := ctx.Param("id")
	id, _ := strconv.Atoi(sId)
	err := modelUser.Delete(id)
	response(ctx, err)
}

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	var data models.User
	errBind := ctx.ShouldBindJSON(&data)
	if errBind != nil {
		response(ctx, errBind)
		log.Fatalln("errBind:", errBind)
		return
	}

	dbUser, err := modelUser.SelectOne(&data)
	if err != nil {
		response(ctx, err)
		return
	}

	if data.Password != dbUser.Password {
		// TODO 输入错误7次锁定24小时
		// TODO 每3个月提示改密码
		response(ctx, errors.New("用户名或密码错误"))
		_, err := modelUser.UpdateUserUuid(dbUser.Id, true)
		if err != nil {
			return
		}
	} else {
		token, _ := modelUser.UpdateUserUuid(dbUser.Id, false)
		responseData(ctx, err, map[string]interface{}{"id": dbUser.Id, "username": dbUser.Username, "nickname": dbUser.Nickname, "email": dbUser.Email, "token": token})
	}
}

// UserLogout 用户登出
func UserLogout(ctx *gin.Context) {
	sId := ctx.Query("id")
	id, _ := strconv.Atoi(sId)
	_, err := modelUser.UpdateUserUuid(id, true)
	response(ctx, err)
}
