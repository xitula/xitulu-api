package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	t "xitulu/types"
)

/*
@Description 默认接口返回
*/
func response(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusOK, t.Res{Code: 1, Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, t.Res{Code: 0, Message: "ok"})
	}
}

/*
@Description 带数据的默认接口返回
*/
func responseData(c *gin.Context, err error, data interface{}) {
	if err != nil {
		c.JSON(http.StatusOK, t.Res{Code: 1, Message: err.Error(), Data: err})
	} else {
		c.JSON(http.StatusOK, t.Res{Code: 0, Message: "ok", Data: data})
	}
}
