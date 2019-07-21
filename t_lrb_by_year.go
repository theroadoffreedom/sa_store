package store

import (
	models "github.com/theroadoffreedom/sa_xorm_model"
)


func QueryTLrbYByStockId(id string) ([]models.TLrbByYear, error) {
	db, _ := GetDB()
	data := make([]models.TLrbByYear, 0)
	err := db.Where("id = ?", id).Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}


func QueryTLrbByYearLastest(id string) (*models.TLrbByYear, error) {
	db, _ := GetDB()
	model := new(models.TLrbByYear)
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

func QueryTLrbByYearList() ([]models.TLrbByYear, error) {
	db, _ := GetDB()
	data := make([]models.TLrbByYear, 0)
	err := db.Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func InsertTLrbByYear(model *models.TLrbByYear) (uint,error) {

	db, _ := GetDB()

	effect_row, err := db.InsertOne(model)
	if err != nil {
		return 0,err
	}
	return uint(effect_row),nil
}

func CountTLrbByYear() (int64, error) {
	db, _ := GetDB()
	model := new(models.TLrbByYear)
	c, err := db.Count(model)
	if err != nil {
		return 0, err
	}
	return c, nil
}


func UpdateTLrbByYear(model *models.TLrbByYear) (int64, error) {
	db, _ := GetDB()
	return db.Where("id = ?", model.Id).And("data_time = ?",model.DataTime).Update(model)
}
