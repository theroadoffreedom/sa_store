package store

type FinanceReportType int
const (
	_ FinanceReportType = iota
	BalanceSheet 			// 资产负债表    
	CashFlowStatement		// 现金流表
	IncomeStatement 			// 利润表
)

type FinanceTimeType int
const (
	_ FinanceTimeType = iota
	Quarter			// 季度 
	Yearly			// 年度
)
