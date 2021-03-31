package store

import (
	"errors"

	models "github.com/theroadoffreedom/sa_xorm_model"
)

const (
	STORE_ERR_CODE_CHECK_TIME_FOUND_EMPTY = "cb check time get empty"
)

func QueryTStock() ([]models.TStock, error) {
	db, _ := GetDB()
	data := make([]models.TStock, 0)
	err := db.Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func FuzzyQueryTStock(keyword string) ([]models.TStock, error) {
	data := make([]models.TStock, 0)
	if len(keyword) == 0 {
		return data, nil
	}

	db, _ := GetDB()
	sql := "select * from t_stock where id like '%" + keyword + "%' or cn like '%" + keyword + "%' or full_name like '%" + keyword + "%'"
	err := db.SQL(sql).Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetAStockCount() (int64, error) {
	db, _ := GetDB()

	stock := new(models.TStock)
	c, err := db.Where("1").Count(stock)
	if err != nil {
		return 0, err
	}
	return c, nil
}

func InsertTStock(model *models.TStock) (uint, error) {
	db, _ := GetDB()

	effect_row, err := db.InsertOne(model)
	if err != nil {
		return 0, err
	}
	return uint(effect_row), nil
}

func GetStockByCbCheckTime() (*models.TStock, error) {
	db, _ := GetDB()
	stock := new(models.TStock)
	has, err := db.Asc("cb_check_time").Get(stock)
	if err != nil {
		return nil, err
	}
	if has {
		return stock, nil
	} else {
		return nil, errors.New(STORE_ERR_CODE_CHECK_TIME_FOUND_EMPTY)
	}
}

func UpdateTStock(model *models.TStock) (int64, error) {
	db, _ := GetDB()
	return db.Id(model.Id).Update(model)
}

func GetStockInfoOrderByInfoCheckTime() (*models.TStock, error) {
	db, _ := GetDB()
	stock := new(models.TStock)
	has, err := db.Asc("info_check_time").Get(stock)
	if err != nil {
		return nil, err
	}
	if has {
		return stock, nil
	} else {
		return nil, errors.New(STORE_ERR_CODE_CHECK_TIME_FOUND_EMPTY)
	}
}
