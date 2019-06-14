package store

import (
	"errors"

	models "github.com/theroadoffreedom/sa_xorm_model"
)

func QuerySzseCsrcMarketPeInfo(t int64) (*models.TCsrcMarketSzSzsePe, error) {
	db, _ := GetDB()
	m := new(models.TCsrcMarketSzSzsePe)
	has, err := db.Where("data_time=?", t).Get(m)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return m, nil
}

func UpdateSzseCsrcMarketPeInfo(model *models.TCsrcMarketSzSzsePe) (uint,error) {
	db, _ := GetDB()
	if model == nil {
		return 0,errors.New(STORE_ERR_CODE_MODEL_EMPTY)
	}

	affected, err := db.Update(model)
	if err != nil {
		return 0,err
	}
	return uint(affected),nil
}

func InsertSzseCsrcMarketPeInfo(model *models.TCsrcMarketSzSzsePe) (uint,error) {
	db, _ := GetDB()
	if model == nil {
		return 0,errors.New(STORE_ERR_CODE_MODEL_EMPTY)
	}

	affected, err := db.InsertOne(model)
	if err != nil {
		return 0,err
	}
	return uint(affected),nil
}
