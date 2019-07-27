package store

import (

	"fmt"
	"errors"

	models "github.com/theroadoffreedom/sa_xorm_model"
	utils "github.com/theroadoffreedom/utils"
)

// return
// - bool, should run,
// - error
func InsertOperatorRunStatisIfNotExistAndUpdateCheckTime(stock_id string, operator_id string) (bool, error) {

	//step1 : check whether exist
	db, _ := GetDB()
	items := make([]models.TAstockOperatorRunStatis,0)
	err := db.Where("stock_id = ? and operator_id = ?", stock_id, operator_id).Find(&items)
	if err != nil {
		return false,err
	}


	// step2: if not exist , insert
	if len(items) == 0 {
		obj := &models.TAstockOperatorRunStatis{
			StockId:stock_id,
			OperatorId:operator_id,
			CheckTime:int(utils.GetCurrentTimestamp()),
			ShouldRun:"yes"}
		eff, err := db.InsertOne(obj)
		if err != nil {
			return false, err
		}
		if eff == 0 {
			return false, errors.New("insert new one error")
		}
		return true, nil
	}

	// step3: if exist, check whether should run
	if len(items) != 1{
		return false, errors.New("operator key error")
	}
	if items[0].ShouldRun == "no" {
		return false,nil
	}

	// step4: if should run , update check time
	items[0].CheckTime = int(utils.GetCurrentTimestamp())
	fmt.Println(items[0].CheckTime)
	eff, err := db.Cols("check_time").Update(&items[0])
	if err != nil {
		return false, err
	}
	if eff == 0 {
		return false, errors.New("update check time error")
	}
	return true,nil
}

func UpdateOperatorRunState(state string, stock_id string, operator_id string) (int64,error) {
	db, _ := GetDB()
	obj := &models.TAstockOperatorRunStatis{
		StockId:stock_id,
		OperatorId:operator_id,
		CheckTime:int(utils.GetCurrentTimestamp()),
		ShouldRun:state}
	return db.Cols("should_run").Update(obj)
}
