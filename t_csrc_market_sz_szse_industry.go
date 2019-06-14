package store

import (
	"errors"

	models "github.com/theroadoffreedom/sa_xorm_model"
)

func QuerySzseCsrcMarketIndustryInfo(t int64) (*models.TCsrcMarketSzSzseIndustry, error) {
	db, _ := GetDB()
	m := new(models.TCsrcMarketSzSzseIndustry)
	has, err := db.Where("data_time=?", t).Get(m)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return m, nil
}

func UpdateSzseCsrcMarketIndustryInfo(model *models.TCsrcMarketSzSzseIndustry) (uint,error) {
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

func InsertSzseCsrcMarketIndustryInfo(model *models.TCsrcMarketSzSzseIndustry) (uint,error) {
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
