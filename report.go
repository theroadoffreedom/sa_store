package store

import (

	"fmt"
	"errors"

	models "github.com/theroadoffreedom/sa_xorm_model"
	utils "github.com/theroadoffreedom/utils"
)

type ReportType struct {
	StockId string			// id
	DataTime int64			// 报表时间
	TimeType FinanceTimeType	// 报表类型，季，年
}

func GetReportDateRange(stock_id string, timeType FinanceTimeType, begin int64,end int64) ([]ReportType, error) {

	if begin == end {
		return nil, errors.New("begin is equal end")
	}
	if begin > end {
		return nil, errors.New("begin is lager than end")
	}
	if len(stock_id) == 0 {
		return nil, errors.New("stock id is empty")
	}

	db, _ := GetDB()
	dateRange := make([]ReportType,0)
	if timeType == Quarter {
		reports := make([]models.TLrbByQuarter,0)
		err := db.Where("id = ?",stock_id).And("data_time >= ?",begin).And("data_time <= ?", end).Find(&reports)
		if err != nil {
			return nil, err
		}
		if len(reports) != 0 {
			for _, r:= range reports {
				obj:= ReportType{
					StockId:stock_id,
					DataTime:int64(r.DataTime),
					TimeType:Quarter}
				dateRange = append(dateRange, obj)
			}
		}
		return dateRange, nil
	}

	if timeType == Yearly {
		reports := make([]models.TLrbByYear,0)
		err := db.Where("id = ?",stock_id).And("data_time >= ?",begin).And("data_time <= ?", end).Find(&reports)
		if err != nil {
			return nil, err
		}
		if len(reports) != 0 {
			for _, r:= range reports {
				obj:= ReportType{
					StockId:stock_id,
					DataTime:int64(r.DataTime),
					TimeType:Yearly}
				dateRange = append(dateRange, obj)
			}
		}
		return dateRange, nil
	}

	// all
	q_reports := make([]models.TLrbByQuarter,0)
	err := db.Where("id= ? ",stock_id).And("data_time >= ?",begin).And("data_time <= ?", end).Find(&q_reports)
	if err != nil {
		return nil, err
	}
	if len(q_reports) != 0 {
		for _, r:= range q_reports {
			obj:= ReportType{
				StockId:stock_id,
				DataTime:int64(r.DataTime),
				TimeType:Quarter}
			dateRange = append(dateRange, obj)
		}
	}

	y_reports := make([]models.TLrbByYear,0)
	err = db.Where("id= ? ",stock_id).And("data_time >= ?",begin).And("data_time <= ?", end).Find(&y_reports)
	if err != nil {
		return nil, err
	}
	if len(y_reports) != 0 {
		for _, r:= range y_reports {
			obj:= ReportType{
				StockId:stock_id,
				DataTime:int64(r.DataTime),
				TimeType:Yearly}
			dateRange = append(dateRange, obj)
		}
	}
	return dateRange, nil
}

func CreateReportId(tType FinanceTimeType, id string,rType FinanceReportType, dataTime int64) string {
	return fmt.Sprintf(
		"%s_%s_%s_%s_%s_%s", 
		GetReportTypeStr(rType), GetTimeTypeReportStr(tType), id,
		utils.GetYearFromTimestamp(dataTime), 
		utils.GetMonthFromTimestamp(dataTime), 
		utils.GetDayFromTimestamp(dataTime))
}

func GetReportId(reportType ReportType) ([]string) {
	ids := make([]string,3)
	ids[0] = CreateReportId(reportType.TimeType,reportType.StockId, BalanceSheet, reportType.DataTime)
	ids[1] = CreateReportId(reportType.TimeType,reportType.StockId, CashStatement, reportType.DataTime)
	ids[2] = CreateReportId(reportType.TimeType,reportType.StockId, ProfitStatement, reportType.DataTime)
	return ids
}

// 
func GetReportData(reportId string) ([]models.TAstockFinanceReportData, error) {
	db, _ := GetDB()
	data := make([]models.TAstockFinanceReportData, 0)
	err := db.Where("report_id = ?", reportId).Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
