package store

import (


	models "github.com/theroadoffreedom/sa_xorm_model"
	utils "github.com/theroadoffreedom/utils"
)


func CountReportData(reportId string)(int64,error) {
		db, _ := GetDB()
		model := models.TAstockFinanceReportData{}
		if len(reportId) != 0 {
			model.ReportId = reportId
		}
		return db.Count(&model)
}


func NewReportData(
	data string, 
	itemId string, 
	reportId string, 
	reportType FinanceReportType, 
	reportTimeType FinanceTimeType) *models.TAstockFinanceReportData {
		rd := new(models.TAstockFinanceReportData)
		rd.Data = data
		rd.ItemId = itemId
		rd.ReportId = reportId
		rd.ReportType = int(reportType)
		rd.ReportTimeType = int(reportTimeType)
		rd.CreateTime = int(utils.GetCurrentTimestamp())
		rd.CheckTime = rd.CreateTime
		return rd
}

func InsertReportItemDataWhenNotExist(obj *models.TAstockFinanceReportData) (int64, error) {
	db, _ := GetDB()
	items := make([]models.TAstockFinanceReportData,0)
	err := db.Where("item_id = ? and report_id = ?", obj.ItemId, obj.ReportId).Find(&items)
	if err != nil {
		return 0,err
	}
	if len(items) != 0 {
		return 0, nil
	}
	return db.InsertOne(obj)
}


func QueryReportData(reportId string, offset uint64, limit uint64)([]models.TAstockFinanceReportData, error) {
	db, _ := GetDB()
	items := make([]models.TAstockFinanceReportData,0)
	if len(reportId) != 0 {
		err := db.Where("report_id = ?",reportId).Limit(int(limit), int(offset)).Find(&items)
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
