package store

import (
	"errors"

	models "github.com/theroadoffreedom/sa_xorm_model"
)



// - bool, true is insert and need init, false is exist and inited
// - error, error 
func DayKLineInsertIfNotExist(id string) (bool, error) {
	db, _ := GetDB()
	items := make([]models.TAstockKLineManage,0)
	err := db.Where("stock_id = ?", id).Find(&items)
	if err != nil {
		return false,err
	}

	c := len(items)
	if c == 0 {
		// inert
		obj := &models.TAstockKLineManage{
			StockId:id,
			DayKLineInit:0}
			eff, err := db.InsertOne(obj)
			if err != nil {
				return false, err
			}
			if eff == 0 {
				return false,errors.New("inert error")
			}
			return true, nil
	}

	if c != 1 {
		return false, errors.New("get models.TAstockKLineManage error, count is not right")
	}

	// c == 1
	if items[0].DayKLineInit == 0 {
		return true,nil
	} else {
		return false,nil
	}
}

func UpdateDayKLineInit(id string, init int) error {
	db, _ := GetDB()
	obj := &models.TAstockKLineManage{
		StockId:id,
		DayKLineInit:init}

	eff, err := db.Cols("day_k_line_init").Update(obj)
	if err != nil {
		return err
	}
	if eff != 1 {
		return errors.New("update error")
	}
	return nil
}
