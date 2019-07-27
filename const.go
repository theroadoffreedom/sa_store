package store

import (
	"errors"
	"strings"

	models "github.com/theroadoffreedom/sa_xorm_model"
)

type FinanceReportType int

const (
	AllSheet        FinanceReportType = iota
	BalanceSheet                      // 资产负债表
	CashStatement                     // 现金流表
	ProfitStatement                   // 利润表
)

type FinanceTimeType int

const (
	AllTime FinanceTimeType = iota
	Quarter                 // 季度
	Yearly                  // 年度
)

type ReportState int

const (
	AllReportState ReportState = iota
	ReportNormal
	ReportInvalid
	ReportFailed
)

func checkReportType(reportType int) error {
	_type := FinanceReportType(reportType)
	if _type == BalanceSheet || _type == CashStatement || _type == ProfitStatement {
		return nil
	}
	return errors.New(STORE_REPORT_TYPE_ERROR)
}

func checkReportTimeType(reportTimeType int) error {
	_type := FinanceTimeType(reportTimeType)
	if _type == Quarter || _type == Yearly {
		return nil
	}
	return errors.New(STORE_REPORT_TIME_TYPE_ERROR)
}

func checkReportState(state int) error {
	_state := ReportState(state)
	if _state == ReportNormal || _state == ReportInvalid || _state == ReportFailed {
		return nil
	}
	return errors.New(STORE_REPORT_STATE_ERROR)
}

func checkIndex(obj *models.TAutoFinanceReportIndex) error {
	// check report type
	err := checkReportType(obj.ReportType)
	if err != nil {
		return err
	}
	err = checkReportTimeType(obj.ReportTimeType)
	if err != nil {
		return err
	}
	err = checkReportState(obj.State)
	if err != nil {
		return err
	}

	return nil
}

func GetReportTypeStoreType(reportType string) FinanceReportType {
        if reportType == "balance" {
                return BalanceSheet
        }
        if reportType == "cash" {
                return CashStatement
        }
        if reportType == "profit" {
                return ProfitStatement
        }
        return AllSheet
}

func GetReportTypeStr(reportType FinanceReportType) string {
	if reportType == BalanceSheet {
		return "balance"
	}
	if reportType == CashStatement {
		return "cash"
	}

	if reportType == ProfitStatement {
		return "profix"
	}
	return "all"
}

func GetTimeTypeStoreType(timeType string) FinanceTimeType {
        if timeType == "quarter" {
                return Quarter
        }
        if timeType == "year" {
                return Yearly
        }
        return AllTime
}

func GetTimeTypeStr(timeType FinanceTimeType) string {
	if timeType == Quarter {
		return "quarter"
	}
	if timeType == Yearly {
		return "year"
	}
	return "all"
}

func GetReportTypeFromId(id string) (string, error) {
	if len(id) == 0 {
		return errors.New("report id error")
	}

	strs := strings.Split(id,"_")
	switch strs[0] {
		case "cash","balance","profit":{
			return strs[0],nil
		}
		default:{
			return "",errors.New("report id is error, parse error")
		}
	}
}

func GetReportTimeTypeFromId(id string) (string, error) {
	if len(id) == 0 {
		return errors.New("report id error")
	}

	strs := strings.Split(id,"_")
	switch strs[0] {
		case "y":{
			return "year",nil
		}
		case "q":{
			return "querter",nil
		}
		default:{
			return "",errors.New("report id is error, parse error")
		}
	}
}
