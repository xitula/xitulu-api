package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"xitulu/models"
	t "xitulu/types"
)

var modelCauserie models.Causerie

func init() {
	modelCauserie = models.Causerie{}
}

// CauseriesAll 获取全部随感
func CauseriesAll(c *gin.Context) {
	var page t.Pagination
	page.CurrentPage, _ = strconv.Atoi(c.Query("currentPage"))
	page.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	if page.CurrentPage != 0 && page.PageSize != 0 {
		data, count, errModel := modelCauserie.SelectAllPage(&page)
		responseData(c, errModel, map[string]interface{}{"list": data, "total": count})
	} else {
		data, errModel := modelCauserie.SelectAll()
		responseData(c, errModel, data)
	}
}

// CauseriesAdd 新增随感
func CauseriesAdd(c *gin.Context) {
	var causerie models.Causerie
	errBind := c.ShouldBindJSON(&causerie)
	if errBind != nil {
		response(c, errBind)
		return
	}
	causerie.CreateDate = time.Now()
	causerie.Status = 1
	errInsert := modelCauserie.Insert(&causerie)
	response(c, errInsert)
}

// CauseriesUpdate 更新随感
func CauseriesUpdate(c *gin.Context) {
	var causerie models.Causerie
	errBind := c.ShouldBindJSON(&causerie)
	if errBind != nil {
		response(c, errBind)
	}
	errUpdate := modelCauserie.Update(causerie.Id, causerie.Content)
	response(c, errUpdate)
}

// CauseriesDelete 删除随感
func CauseriesDelete(c *gin.Context) {
	sId := c.Param("id")
	id, _ := strconv.Atoi(sId)
	errUpdate := modelCauserie.Delete(id)
	response(c, errUpdate)
}
