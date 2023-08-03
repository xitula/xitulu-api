package model

import (
	"errors"
	"log"
	t "xitulu/types"
	u "xitulu/util"
)

// 查询所有随感数据
func SelectCauseriesAll() (*[]t.Causerie, error) {
	var causeries []t.Causerie
	result := orm.Table("causeries").Where("status = 1").Order("create_date DESC").Find(&causeries)
	err := result.Error
	if err != nil {
		log.Fatalln("SelectCauseriesAllError:", err)
		return nil, err
	}

	return &causeries, nil
}

// 分页查询所有随感数据
func SelectCauseriesPage(page *t.Pagination) (*map[string]interface{}, error) {
	var causeries []t.Causerie
	var count int64
	result := orm.Table("causeries").Where("status = 1").Count(&count).Order("create_date DESC").Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).Find(&causeries)
	err := result.Error
	if err != nil {
		log.Fatalln("SelectCauseriesAllError:", err)
		return nil, err
	}
	finalResult := map[string]interface{}{
		"list":  causeries,
		"count": count,
	}
	return &finalResult, nil
}

// 插入随感数据
func InsertCauserie(causerie *t.Causerie) error {
	causerie.Status = 1
	causerie.CreateDate = u.GetMysqlNow()
	causerie.LastUpdateDate = u.GetMysqlNow()
	result := orm.Create(causerie)
	errMod := result.Error
	if errMod != nil {
		log.Fatalln("InsertCauserieError:", errMod)
	}
	return nil
}

// 依据ID更新随感数据
func UpdateCauserie(id int, content string) error {
	lastUpdateDate := u.GetMysqlNow()
	result := orm.Table("causeries").Where("id = ?", id).Updates(t.Causerie{Id: id, Content: content, LastUpdateDate: lastUpdateDate})
	errUpd := result.Error
	if errUpd != nil {
		log.Fatalln("UpdateCauserieError:", errUpd)
	}
	if result.RowsAffected == 0 {
		return errors.New("id错误")
	}

	return nil
}

// 依据ID删除随感
func DeleteCauserie(id int) error {
	lastUpdateDate := u.GetMysqlNow()
	result := orm.Table("causeries").Where("id = ?", id).Updates(map[string]interface{}{"id": id, "status": 0, "last_update_date": lastUpdateDate})
	errUpd := result.Error
	if errUpd != nil {
		log.Fatalln("UpdateCauserieError:", errUpd)
	}
	if result.RowsAffected == 0 {
		return errors.New("id错误")
	}

	return nil
}
