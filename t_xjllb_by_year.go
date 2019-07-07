package store

import (
	models "github.com/theroadoffreedom/sa_xorm_model"
)

func QueryTXjllbYByStockId(id string) ([]models.TXjllbByYear, error) {
	db, _ := GetDB()
	data := make([]models.TXjllbByYear, 0)
	err := db.Where("id = ?", id).Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}


func QueryTXjllbByYearLastest(id string) (*models.TXjllbByYear, error) {
	db, _ := GetDB()
	model := new(models.TXjllbByYear)
	has, err := db.Where("id = ?", id).Asc("data_time").Get(model)
	if err != nil {
		return nil, err
	}
	if has {
		return model, nil
	} else {
		return nil, nil
	}

}

func QueryTXjllbByYearList() ([]models.TXjllbByYear, error) {
	db, _ := GetDB()
	data := make([]models.TXjllbByYear, 0)
	err := db.Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func InsertTXjllbByYear(model *models.TXjllbByYear) (uint,error) {
	db, _ := GetDB()

	effect_row, err := db.InsertOne(model)
	if err != nil {
		return 0,err
	}
	return uint(effect_row),nil
}

func CountTXjllbByYear() (int64, error) {
	db, _ := GetDB()
	model := new(models.TXjllbByYear)
	c, err := db.Count(model)
	if err != nil {
		return 0, err
	}
	return c, nil
}
