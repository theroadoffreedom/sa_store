package store

import (
	"testing"
	"time"

	models "github.com/theroadoffreedom/sa_xorm_model"
)

const (
	TEST_DB_NAME="china_stock_data_db"
	TEST_DB_USER = "root"
	TEST_DB_USER_PW = "busysa"
	TEST_DB_IP = "192.168.3.7"
	TEST_DB_PORT  = 4000
	TEST_STOCK_ID = "test"
	TEST_TM = 1405544146
)


func TestDayKLineIsExist(t *testing.T) {
	// 
	err := InitStore(TEST_DB_IP,TEST_DB_PORT,TEST_DB_USER, TEST_DB_USER_PW, TEST_DB_NAME)
	if err != nil {
		t.Error(err)
		return
	}

	tm := time.Now()
	exist,err := DayKLineIsExist(TEST_STOCK_ID, tm.Unix())
	if err != nil {
		t.Error(err.Error())
		return
	}
	if exist {
		t.Error("should be not exist")
		return
	}
	d := models.TAstockDayKLineData{
		StockId:TEST_STOCK_ID,
		Day:tm,
		Open:"0",
		High:"0",
		Low:"0",
		Volumn:0,
		Close:"0"}
	err = InsertTAstockDayKLineData(d)
	if err != nil {
		t.Error(err)
		return
	}
	exist,err = DayKLineIsExist(TEST_STOCK_ID, tm.Unix())
	if err != nil {
		t.Error(err.Error())
		return
	}
	if !exist {
		t.Error("should be exist")
		return
	}
}
