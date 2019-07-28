package store

import (
	"errors"

	models "github.com/theroadoffreedom/sa_xorm_model"
	utils "github.com/theroadoffreedom/utils"
)

func DayKLineIsExist(id string, t int64) (bool,error) {
	db, _ := GetDB()
	items := make([]models.TAstockDayKLineData,0)
	err := db.Where("stock_id = ? and day = ?", id, utils.TimeFormatDBString(t)).Find(&items)
	if err != nil {
		return false,err
	}
	if len(items) == 0 {
		return false,nil
	}
	return true,nil
}

func InsertTAstockDayKLineData(d models.TAstockDayKLineData) error {
	db, _ := GetDB()
	eff, err := db.InsertOne(&d)
	if err != nil{
		return err
	}
	if eff == 0{
		return errors.New("inert error")
	}
	return nil
}
