package model

import (
	"log"
	t "xitulu/types"
	u "xitulu/util"
)

func SelectCauseriesAll() (*[]t.Causerie, error) {
	var causeries []t.Causerie
	var count int64
	result := orm.Table("causeries").Order("create_date DESC").Find(&causeries).Count(&count)
	err := result.Error
	if err != nil {
		log.Fatalln("SelectCauseriesAllError:", err)
		return nil, err
	}

	return &causeries, nil
}

func SelectCauseriesPage(page *t.Pagination) (*map[string]interface{}, error) {
	var causeries []t.Causerie
	var count int64
	result := orm.Table("causeries").Count(&count).Order("create_date DESC").Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).Find(&causeries)
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

func InsertCauserie(causerie *t.Causerie) error {
	causerie.Status = 1
	causerie.CreateDate = u.GetMysqlNow()
	// log.Printf("causerie: %+v\n", causerie)
	result := orm.Create(causerie)
	errMod := result.Error
	if errMod != nil {
		log.Fatalln("InsertCauserieError:", errMod)
	}
	return nil
}
