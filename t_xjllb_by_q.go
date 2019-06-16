package store

import (
	"strconv"
	"fmt"
	"time"

	models "github.com/theroadoffreedom/sa_xorm_model"
	utils "github.com/theroadoffreedom/utils"
)

func QueryTXjllbByQLastestCount(id string) (int64, error) {
	db, _ := GetDB()
	model := new(models.TXjllbByQuarter)

	year := utils.GetCurrentYear()
	next_year, _ := strconv.Atoi(year)
	next_year = next_year + 1
	beginTimestamp, _ := utils.Translate2Timestamp(fmt.Sprintf("%s-01-01T00:00:00+08:00", year), time.RFC3339)
	endTimestamp, _ := utils.Translate2Timestamp(fmt.Sprintf("%d-01-01T00:00:00+08:00", next_year), time.RFC3339)

	model.Id = id
	return db.Where(" data_time >? and data_time < ?", beginTimestamp, endTimestamp).Count(model)
}

func QueryTXjllbByQuarterList() ([]models.TXjllbByQuarter, error) {
	db, _ := GetDB()
	data := make([]models.TXjllbByQuarter, 0)
	err := db.Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func InsertTXjllbByQuarter(model *models.TXjllbByQuarter) (uint,error) {

	db, _ := GetDB()

	effect_row, err := db.InsertOne(model)
	if err != nil {
		return 0,err
	}
	return uint(effect_row),nil
}

func CheckTXjllbByQuarter(date int64, id string) (bool, error) {

	db, _ := GetDB()
	model := new(models.TXjllbByQuarter)
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

func CountTXjllbByQuarter() (int64, error) {
	db, _ := GetDB()
	model := new(models.TXjllbByQuarter)
	c, err := db.Count(model)
	if err != nil {
		return 0, err
	}
	return c, nil
}
