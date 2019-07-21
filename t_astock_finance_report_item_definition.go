package store

import (

	"fmt"
	"errors"

	models "github.com/theroadoffreedom/sa_xorm_model"
	utils "github.com/theroadoffreedom/utils"
)

func CountReportItem(itemType string)(int64,error) {
		db, _ := GetDB()
		model := models.TAstockFinanceReportItemDefinition{}
		if len(itemType) != 0 {
			model.ItemType = itemType
		}
		return db.Count(&model)
}


func NewReportItem(cn string, unit string, itemType string) (*models.TAstockFinanceReportItemDefinition,error) {
	item := new(models.TAstockFinanceReportItemDefinition)
	item.Cn = cn
	item.Unit = unit
	item.ItemType = itemType
	item.CreateTime = int(utils.GetCurrentTimestamp()) 
	// balance_item_num
	c, err := CountReportItem(itemType)
	if err != nil {
		return nil, err
	}
	item.Id = fmt.Sprintf("%s_item_%c", itemType, c)
	return item, nil
}


func FuzzyQueryReportItem(
	keyword string,
	offset uint64, 
	limit uint64) ([]models.TAstockFinanceReportItemDefinition, error) {

	data := make([]models.TAstockFinanceReportItemDefinition, 0)
	if len(keyword)== 0 {
		return data,nil	
	}

	db, _ := GetDB()
	sql := "select * from t_astock_finance_report_item_definition where cn like '%" + keyword + "%'" 
	err := db.SQL(sql).Limit(int(limit), int(offset)).Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func QueryReportItem(id string, cn string, itemType string, offset uint64, limit uint64) ([]models.TAstockFinanceReportItemDefinition, error){
	items := make([]models.TAstockFinanceReportItemDefinition,0)
	db, _ := GetDB()

	// just use id
	if len(id) != 0 {
		err := db.Id(id).Find(&items)
		if err != nil {
			return nil, err
		}
		return items, nil
	}

	// just cn
	if len(id) != 0 {
		err := db.Where("cn = ?",cn).Find(&items)
		if err != nil {
			return nil, err
		}
		return items, nil
	}

	// check
	if limit == 0 {
		return nil, errors.New(STORE_LIMIT_IS_ZERO_ERROR)
	}

	if len(itemType) != 0 {
		err := db.Where("item_type = ?",itemType).Limit(int(limit), int(offset)).Find(&items)
		if err != nil {
			return nil,err 
		}
		return items, nil
	} else {
		err := db.Limit(int(limit), int(offset)).Find(&items)
		if err != nil {
			return nil,err 
		}
		return items, nil
	}
}

func InsertReportItemWhenNotExist(obj *models.TAstockFinanceReportItemDefinition) (string,int64, error) {
	db, _ := GetDB()
	items := make([]models.TAstockFinanceReportItemDefinition,0)
	err := db.Where("cn = ? and unit = ? and item_type = ?", obj.Cn, obj.Unit, obj.ItemType).Find(&items)
	if err != nil {
		return "",0,err
	}
	if len(items) != 0 {
		if len(items)==1 {
			return items[0].Id, 0, nil
		} else {
			return "",0,errors.New("data has too many version, error! error!")
		}
	}
	eff, err := db.InsertOne(obj)
	return obj.Id, eff,err
}
