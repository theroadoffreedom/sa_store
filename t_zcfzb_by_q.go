package store

import (
	"fmt"
	"strconv"
	"time"

	models "github.com/theroadoffreedom/sa_xorm_model"
	utils "github.com/theroadoffreedom/utils"
)

func QueryTZcfzbQByStockId(id string) ([]models.TZcfzbByQuarter, error) {
	db, _ := GetDB()
	data := make([]models.TZcfzbByQuarter, 0)
	err := db.Where("id = ?", id).Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}


func QueryTZcfzbByQLastestCount(id string) (int64, error) {
	db, _ := GetDB()
	model := new(models.TZcfzbByQuarter)

	year := utils.GetCurrentYear()
	next_year, _ := strconv.Atoi(year)
	next_year = next_year + 1
	beginTimestamp, _ := utils.Translate2Timestamp(fmt.Sprintf("%s-01-01T00:00:00+08:00", year), time.RFC3339)
	endTimestamp, _ := utils.Translate2Timestamp(fmt.Sprintf("%d-01-01T00:00:00+08:00", next_year), time.RFC3339)
	model.Id = id
	return db.Where(" data_time >? and data_time < ?", beginTimestamp, endTimestamp).Count(model)
}

func QueryTZcfzbByQuarterList() ([]models.TZcfzbByQuarter, error) {
	db, _ := GetDB()
	data := make([]models.TZcfzbByQuarter, 0)
	err := db.Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func InsertTZcfzbByQuarter(model *models.TZcfzbByQuarter) (uint,error) {

	db, _ := GetDB()

	effect_row, err := db.InsertOne(model)
	if err != nil {
		return 0,err
	}
	return uint(effect_row),nil
}

func CheckTZcfzbByQuarter(date int64, id string) (bool, error) {

	db, _ := GetDB()
	model := new(models.TZcfzbByQuarter)
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

func CountTZcfzbByQuarter() (int64, error) {
	db, _ := GetDB()
	model := new(models.TZcfzbByQuarter)
	c, err := db.Count(model)
	if err != nil {
		return 0, err
	}
	return c, nil
}

func UpdateTZcfzbByQuarter(model *models.TZcfzbByQuarter) (int64, error) {
	db, _ := GetDB()
	return db.Where("id = ?", model.Id).And("data_time = ?",model.DataTime).Update(model)
}
