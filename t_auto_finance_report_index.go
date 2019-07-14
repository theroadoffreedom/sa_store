package store

import (

	"errors"

	models "github.com/theroadoffreedom/sa_xorm_model"
	utils "github.com/theroadoffreedom/utils"
)

func CountReportIndex(
	id string,
	dataTime int,
	state ReportState,
	reportType FinanceReportType,
	reportTimeType FinanceTimeType) (int64, error) {

		db, _ := GetDB()
		model := models.TAutoFinanceReportIndex{}
		if len(id) != 0 {
			model.Id = id
		}
		if dataTime != 0 {
			model.DataTime = dataTime
		}
		if state != AllReportState {
			model.State = int(state)
		}
		if reportType != AllSheet {
			model.ReportType = int(reportType)
		}
		if reportTimeType != AllTime {
			model.ReportTimeType = int(reportTimeType)

		}
		c, err := db.Count(&model)
		if err != nil {
			return 0, err
		}
		return c, nil
}


// report id format : [balance|cash|profix]_[stockid]_[y|q]_[data_time]
func NewReportIndex(
	id string, 
	dataTime int, 
	state ReportState, 
	reportType FinanceReportType, 
	reportTimeType FinanceTimeType) *models.TAutoFinanceReportIndex {

		if len(id) == 0 {
			return nil
		}

		inx := new(models.TAutoFinanceReportIndex)
		inx.ReportId = fmt.Sprintf("%s_%s_%s_%d", GetReportTypeStr(reportType), id, GetTimeTypeStr(reportTimeType), dataTime)
		inx.Id = id
		inx.DataTime = dataTime
		inx.State = int(state)
		inx.ReportType = int(reportType)
		inx.ReportTimeType = int(reportTimeType)
		inx.CreateTime = int(utils.GetCurrentTimestamp())
		inx.CheckTime =inx.CreateTime
		return inx
}

// get
func QueryReportIndex(
	id string, reportType FinanceReportType, 
	reportTimeType FinanceTimeType, 
	offset uint64, limit uint64) ([]models.TAutoFinanceReportIndex, error){

		// check
		if limit == 0 {
			return nil, errors.New(STORE_LIMIT_IS_ZERO_ERROR)
		}

		indexs := make([]models.TAutoFinanceReportIndex,0)
		db, _ := GetDB()

		if reportTimeType == AllTime && reportType == AllSheet {
			if len(id) != 0 {
				err := db.Where("id = ?",id).Limit(int(limit), int(offset)).Find(&indexs)
				return indexs, err
			} else {
				err := db.Limit(int(limit), int(offset)).Find(&indexs)
				return indexs, err
			}
		}

		if reportTimeType != AllTime && reportType == AllSheet {
			err := db.Where("id = ? AND report_time_type = ?",id,int(reportTimeType)).Limit(int(limit), int(offset)).Find(&indexs)
			return indexs, err
		}

		if reportTimeType == AllTime && reportType != AllSheet {
			err := db.Where("id = ? AND report_type = ?",id,int(reportType)).Limit(int(limit), int(offset)).Find(&indexs)
			return indexs, err
		}

		if len(id) != 0 {
			err := db.Where("id = ? AND report_type = ? AND report_time_type = ?",id,int(reportType),int(reportTimeType)).Limit(int(limit), int(offset)).Find(&indexs)
			return indexs,err
		} else {
			err := db.Where("report_type = ? AND report_time_type = ?",int(reportType),int(reportTimeType)).Limit(int(limit), int(offset)).Find(&indexs)
			return indexs,err
		}
}


// err == nil
// - affect == 0, the obj is exist
// - affect == 1, the obj is not exist and insert success
// err != nil
// - error
func InsertReportIndexWhenNotExist(obj *models.TAutoFinanceReportIndex) (int64,error) {

	err := checkIndex(obj)
	if err != nil {
		return 0,err
	}

	db, _ := GetDB()
	index := make([]models.TAutoFinanceReportIndex,0)
	err = db.Where("id = ? AND data_time = ? AND report_type = ?", obj.Id,obj.DataTime, obj.ReportType).Find(&index)
	if err != nil {
		return 0,err
	}
	if len(index) != 0 {
		return 0,nil
	}

	return db.InsertOne(obj)
}


// delete data!!!!
func DeleteReportIndex(obj *models.TAutoFinanceReportIndex) (int64,error) {
	db, _ := GetDB()
	err := checkIndex(obj)
	if err!= nil {
		return 0,err
	}

	_obj := new(models.TAutoFinanceReportIndex)
	_obj.Id = obj.Id
	_obj.DataTime = obj.DataTime
	_obj.ReportType = obj.ReportType
	return db.Delete(_obj)
}
