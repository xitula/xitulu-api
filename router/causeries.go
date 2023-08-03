package router

import (
	"strconv"
	"xitulu/model"

	t "xitulu/types"

	"github.com/gin-gonic/gin"
)

func registerCauseries(r *gin.Engine) {
	// 查询所有随感数据
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

	// 新增随感
	r.POST("/causeries", func(c *gin.Context) {
		var causerie t.Causerie
		errBind := c.ShouldBindJSON(&causerie)
		if errBind != nil {
			response(c, errBind)
		}
		errInsert := model.InsertCauserie(&causerie)
		response(c, errInsert)
	})

	// 更新随感
	r.PUT("/causeries", func(c *gin.Context) {
		var causerie t.Causerie
		errBind := c.ShouldBindJSON(&causerie)
		if errBind != nil {
			response(c, errBind)
		}
		errUpdate := model.UpdateCauserie(causerie.Id, causerie.Content)
		response(c, errUpdate)
	})

	// 依据ID删除随感
	r.DELETE("/causeries/:id", func(c *gin.Context) {
		sId := c.Param("id")
		id, _ := strconv.Atoi(sId)
		errUpdate := model.DeleteCauserie(id)
		response(c, errUpdate)
	})
}
