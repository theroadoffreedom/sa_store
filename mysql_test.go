package store

import (
	"testing"

	utils "github.com/theroadoffreedom/utils"
)

const (
	TEST_DB_NAME="china_stock_data_db"
	TEST_DB_USER = "root"
	TEST_DB_USER_PW = "busysa"
	TEST_DB_IP = "192.168.3.7"
	TEST_DB_PORT  = 4000
	TEST_STOCK_ID = "000333"
)

func TestQueryQByStockId(t *testing.T) {

	// 
	err := InitStore(TEST_DB_IP,TEST_DB_PORT,TEST_DB_USER, TEST_DB_USER_PW, TEST_DB_NAME)
	if err != nil {
		t.Error(err)
		return
	} 

	data,err := QueryTLrbQByStockId(TEST_STOCK_ID)
	if err != nil {
		t.Error(err)
		return
	}

	if len(data) == 0 {
		t.Error("data is empty")
	}

	t.Log(len(data))
	for _, d := range data {
		t.Log(d.Id,utils.ToHumanString(int64(d.DataTime)))
	}
}

func TestQueryYByStockId(t *testing.T) {

	// 
	err := InitStore(TEST_DB_IP,TEST_DB_PORT,TEST_DB_USER, TEST_DB_USER_PW, TEST_DB_NAME)
	if err != nil {
		t.Error(err)
		return
	} 

	data,err := QueryTLrbYByStockId(TEST_STOCK_ID)
	if err != nil {
		t.Error(err)
		return
	}

	if len(data) == 0 {
		t.Error("data is empty")
	}

	t.Log(len(data))
	for _, d := range data {
		t.Log(d.Id,utils.ToHumanString(int64(d.DataTime)))
	}
}

func TestInsertReportIndexWhenNotExist(t *testing.T) {

	TEST_STOCK_ID := "000000"
	err := InitStore(TEST_DB_IP,TEST_DB_PORT,TEST_DB_USER, TEST_DB_USER_PW, TEST_DB_NAME)
	if err != nil {
		t.Error(err)
		return
	} 

	indexObj := NewReportIndex(TEST_STOCK_ID,int(1562510153), ReportNormal,BalanceSheet,Quarter)
	aff, err := InsertReportIndexWhenNotExist(indexObj)
	if err != nil {
		t.Error(err)
		return
	}
	if aff == 0 {
		t.Error("insert not affect")
		return
	}

	// query for check
	indexs, err := QueryReportIndex(TEST_STOCK_ID,BalanceSheet,Quarter,0,10)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(indexs)
	if len(indexs) != 1 || indexs[0].Id != TEST_STOCK_ID{
		t.Error("query error")
		return 
	}

	t.Log("insert success")

	// insert twice
	aff, err = InsertReportIndexWhenNotExist(indexObj)
	if err != nil {
		t.Error(err)
		return
	}
	if aff != 0 {
		t.Error("insert twice error")
		return
	}

	// delete
	aff, err =  DeleteReportIndex(indexObj)
	if err != nil {
		t.Error(err)
		return
	}
	if aff ==0 {
		t.Error("delete error")
		return
	}
	
	// query for check
	indexs, err = QueryReportIndex(TEST_STOCK_ID,BalanceSheet,Quarter,0,10)
	if err != nil {
		t.Error(err)
		return
	}
	if len(indexs) != 0 {
		t.Error("delete error")
		return 
	}

}


func TestQueryNoId(t *testing.T) {
	// query for no id
	indexs, err := QueryReportIndex("",BalanceSheet,Quarter,0,10)
	if err != nil {
		t.Error(err)
		return
	}
	if len(indexs) == 0 {
		t.Error("query error")
		return 
	}
	t.Log(indexs)
}
