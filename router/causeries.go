package router

import (
	"strconv"
	"xitulu/model"

	t "xitulu/types"

	"github.com/gin-gonic/gin"
)

func registerCauseries(r *gin.Engine) {
	r.GET("/causeries", func(c *gin.Context) {
		var page t.Pagination
		page.CurrentPage, _ = strconv.Atoi(c.Query("currentPage"))
		page.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
		if page.CurrentPage != 0 && page.PageSize != 0 {
			data, errModel := model.SelectCauseriesPage(&page)
			responseData(c, errModel, data)
		} else {
			data, errModel := model.SelectCauseriesAll()
			responseData(c, errModel, data)
		}
	})

	r.POST("/causeries", func(c *gin.Context) {
		var causerie t.Causerie
		errBind := c.ShouldBindJSON(&causerie)
		if errBind != nil {
			response(c, errBind)
		}
		errInsert := model.InsertCauserie(&causerie)
		response(c, errInsert)
	})
}
