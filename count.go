package store

import (
	"errors"
)

func CountTable(t FinanceReportType, time FinanceTimeType) (int64, error) {

	if t == BalanceSheet {
		if time == Quarter {
			return CountTZcfzbByQuarter()	
		} else {
			return CountTZcfzbByYear()
		}
	}

	if t == IncomeStaement {
		if time == Quarter {
			return CountTLrbByQuarter()
		} else {
			return CountTLrbByYear()
		}
	}

	if t == CashFlowStatement {
		if time == Quarter {
			return CountTXjllbByQuarter()
		} else {
			return CountTXjllbByYear()
		}
	}

	return -1, errors.New("unsupport type,time") 
}
