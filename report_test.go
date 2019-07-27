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


func TestGetReportDateRange(t *testing.T) {

	// 
	err := InitStore(TEST_DB_IP,TEST_DB_PORT,TEST_DB_USER, TEST_DB_USER_PW, TEST_DB_NAME)
	if err != nil {
		t.Error(err)
		return
	} 


	data, err:=GetReportDateRange(TEST_STOCK_ID, Yearly, 0, utils.GetCurrentTimestamp())
	if err != nil {
		t.Error(err)
		return
	}

	if len(data)==0 {
		t.Error("data is empty")
		return 
	}

	targatId := ""
	for _, d:=range data {
		ids := GetReportId(d)
		t.Log(ids)
		if targatId == "" {
			targatId = ids[0]
		}
	}

	t.Log("test get report data", targatId)
	rdata,err := GetReportData(targatId)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(rdata)
}
