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

	if t == ProfitStatement {
		if time == Quarter {
			return CountTLrbByQuarter()
		} else {
			return CountTLrbByYear()
		}
	}

	if t == CashStatement {
		if time == Quarter {
			return CountTXjllbByQuarter()
		} else {
			return CountTXjllbByYear()
		}
	}

	return -1, errors.New("unsupport type,time")
}
