package store

import (
	"testing"
)

const (
	TEST_DB_NAME="china_stock_data_db"
	TEST_DB_USER = "root"
	TEST_DB_USER_PW = "busysa"
	TEST_DB_IP = "192.168.3.7"
	TEST_DB_PORT  = 4000
	TEST_STOCK_ID = "test"
)


func TestDayKLineInsertIfNotExist(t *testing.T) {
	// 
	err := InitStore(TEST_DB_IP,TEST_DB_PORT,TEST_DB_USER, TEST_DB_USER_PW, TEST_DB_NAME)
	if err != nil {
		t.Error(err)
		return
	}

	i,err := DayKLineInsertIfNotExist(TEST_STOCK_ID)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !i {
		t.Error("insert failed")
		return
	}

	err = UpdateDayKLineInit(TEST_STOCK_ID,1)
	if err != nil {
		t.Error(err)
		return
	}

	i,err = DayKLineInsertIfNotExist(TEST_STOCK_ID)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if i {
		t.Error("no need to be run")
		return
	}
}
