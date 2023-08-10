package router

import (
	"github.com/gin-gonic/gin"
	"xitulu/router/handler"
)

func registerUsers(r *gin.Engine) {
	// 查询所有用户
	r.GET("/users", handler.UserGetAll)
	// 新增用户
	r.POST("/users", handler.UserAdd)
	// 更新用户数据
	r.PUT("/users", handler.UserUpdate)
	// 删除用户
	r.DELETE("/users/:id", handler.UserDelete)
	// 用户登录
	r.POST("/users/login", handler.UserLogin)
	// 用户登出
	r.GET("/users/logout", handler.UserLogout)
}
