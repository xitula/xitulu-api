package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"xitulu/models"
	"xitulu/types"
)

var modelCauserie models.Causerie

func init() {
	modelCauserie = models.Causerie{}
}

// CauseriesAll 获取全部随感
func CauseriesAll(c *gin.Context) {
	var params types.GetAllParam
	if err := c.ShouldBindQuery(&params); err != nil {
		response(c, err)
		return
	}
	if err := validate.Struct(&params); err != nil {
		response(c, err)
		return
	}
	if params.CurrentPage != 0 && params.PageSize != 0 {
		data, count, errModel := modelCauserie.SelectAllPage(&params)
		responseData(c, errModel, map[string]interface{}{"list": data, "total": count})
	} else {
		data, errModel := modelCauserie.SelectAll()
		responseData(c, errModel, data)
	}
}

// CauseriesAdd 新增随感
func CauseriesAdd(c *gin.Context) {
	// TODO 越权问题
	var causerie models.Causerie
	errBind := c.ShouldBindJSON(&causerie)
	if errBind != nil {
		response(c, errBind)
		return
	}
	if errs := validate.Struct(&causerie); errs != nil {
		response(c, errs)
		return
	}
	causerie.CreateDate = time.Now()
	causerie.Status = 1
	errInsert := modelCauserie.Insert(&causerie)
	response(c, errInsert)
}

// CauseriesUpdate 更新随感
func CauseriesUpdate(c *gin.Context) {
	// TODO 越权问题
	var causerie models.Causerie
	errBind := c.ShouldBindJSON(&causerie)
	if errBind != nil {
		response(c, errBind)
	}
	if errs := validate.Struct(&causerie); errs != nil {
		response(c, errs)
		return
	}
	errUpdate := modelCauserie.Update(causerie.Id, causerie.Content)
	response(c, errUpdate)
}

// CauseriesDelete 删除随感
func CauseriesDelete(c *gin.Context) {
	// TODO 越权问题
	sId := c.Param("id")
	id, _ := strconv.Atoi(sId)
	errUpdate := modelCauserie.Delete(id)
	response(c, errUpdate)
}
