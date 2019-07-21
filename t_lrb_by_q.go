package store

import (
	"fmt"
	"strconv"
	"time"

	models "github.com/theroadoffreedom/sa_xorm_model"
	utils "github.com/theroadoffreedom/utils"
)

func QueryTLrbQByStockId(id string) ([]models.TLrbByQuarter, error) {
	db, _ := GetDB()
	data := make([]models.TLrbByQuarter, 0)
	err := db.Where("id = ?", id).Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func QueryTLrbByQLastestCount(id string) (int64, error) {
	db, _ := GetDB()
	model := new(models.TLrbByQuarter)

	year := utils.GetCurrentYear()
	next_year, _ := strconv.Atoi(year)
	next_year = next_year + 1
	beginTimestamp, _ := utils.Translate2Timestamp(fmt.Sprintf("%s-01-01T00:00:00+08:00", year), time.RFC3339)
	endTimestamp, _ := utils.Translate2Timestamp(fmt.Sprintf("%d-01-01T00:00:00+08:00", next_year), time.RFC3339)

	model.Id = id
	return db.Where(" data_time >? and data_time < ?", beginTimestamp, endTimestamp).Count(model)
}

func QueryTLrbByQuarterList() ([]models.TLrbByQuarter, error) {
	db, _ := GetDB()
	data := make([]models.TLrbByQuarter, 0)
	err := db.Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func InsertTLrbByQuarter(model *models.TLrbByQuarter) (uint,error) {

	db, _ := GetDB()

	effect_row, err := db.InsertOne(model)
	if err != nil {
		return 0,err
	}
	return uint(effect_row),nil
}

func CheckTLrbByQuarter(date int64, id string) (bool, error) {

	db, _ := GetDB()
	model := new(models.TLrbByQuarter)
	model.Id = id
	c, err := db.Where("data_time = ?", date).Count(model)
	if err != nil {
		return false, err
	}
	if c == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func CountTLrbByQuarter() (int64, error) {
	db, _ := GetDB()
	model := new(models.TLrbByQuarter)
	c, err := db.Count(model)
	if err != nil {
		return 0, err
	}
	return c, nil
}


func UpdateTLrbByQuarter(model *models.TLrbByQuarter) (int64, error) {
	db, _ := GetDB()
	return db.Where("id = ?", model.Id).And("data_time = ?",model.DataTime).Update(model)
}
