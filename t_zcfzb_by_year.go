package store

import (
	models "github.com/theroadoffreedom/sa_xorm_model"
)

func QueryTZcfzbYByStockId(id string) ([]models.TZcfzbByYear, error) {
	db, _ := GetDB()
	data := make([]models.TZcfzbByYear, 0)
	err := db.Where("id = ?", id).Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func QueryTZcfzbByYearLastest(id string) (*models.TZcfzbByYear, error) {
	db, _ := GetDB()
	model := new(models.TZcfzbByYear)
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

func QueryTZcfzbByYearList() ([]models.TZcfzbByYear, error) {
	db, _ := GetDB()
	data := make([]models.TZcfzbByYear, 0)
	err := db.Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func InsertTZcfzbByYear(model *models.TZcfzbByYear) (uint,error) {
	db, _ := GetDB()

	effect_row, err := db.InsertOne(model)
	if err != nil {
		return 0,err
	}
	return uint(effect_row),nil
}


func CountTZcfzbByYear() (int64, error) {
	db, _ := GetDB()
	model := new(models.TZcfzbByYear)
	c, err := db.Count(model)
	if err != nil {
		return 0, err
	}
	return c, nil
}

func UpdateTZcfzbByYear(model *models.TZcfzbByYear) (int64, error) {
	db, _ := GetDB()
	return db.Where("id = ?", model.Id).And("data_time = ?",model.DataTime).Update(model)
}
