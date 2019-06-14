package store

import (
	models "github.com/theroadoffreedom/sa_xorm_model"
)

func QueryCsrcList() ([]models.TCsrc, error) {
	db, _ := GetDB()
	data := make([]models.TCsrc, 0)
	err := db.Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UpdateTCsrc(model *models.TCsrc) (uint, error) {
	db, _ := GetDB()

	sql := "update t_csrc set csrc_cn_name = ?,industry_code=?, csrc_id=?"
	res, err := db.Exec(sql, model.CsrcCnName, model.IndustryCode, model.CsrcId)
	if err != nil {
		return 0,err
	}
	effect_row, _ := res.RowsAffected()
	return uint(effect_row),nil
}

func InsertTCsrc(model *models.TCsrc) (uint,error) {
	db, _ := GetDB()

	effect_row, err := db.InsertOne(model)
	if err != nil {
		return 0,err
	}
	return uint(effect_row),nil
}
