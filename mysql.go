package store

import (
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type StoreManger struct {
	sync.Mutex
	dbName     string // db名字
	dbUser     string //登录用户名
	dbIp       string // db ip
	dbPort     int    // db port
	dbPw       string // db password
	inited     bool   // 是否已经初始化
	connnected bool   // 是否连接

	// xorm
	engine *xorm.Engine
}

var storeManager *StoreManger = nil
var once sync.Once

func GetInstance() (*StoreManger, error) {
	once.Do(func() {
		storeManager = &StoreManger{}
		storeManager.inited = false
		storeManager.connnected = false
	})
	return storeManager, nil
}

func (s *StoreManger) Init(ip string, port int, user string, password string, dbname string) error {

	s.dbUser = user
	s.dbPw = password
	s.dbName = dbname
	s.dbIp = ip
	s.dbPort = port
	s.inited = true

	return s.connect()
}

func (s *StoreManger) connect() error {

	if !s.connnected {
		// 数据库名称:数据库连接密码@(数据库地址:3306)/数据库实例名称?charset=utf8
		params := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", s.dbUser, s.dbPw, s.dbIp, s.dbPort, s.dbName)
		engine, err := xorm.NewEngine("mysql", params)
		if err != nil {
			return err
		}

		//engine.Logger().SetLevel(core.LOG_DEBUG)
		s.engine = engine

		err = s.engine.Ping()
		if err != nil {
			return err
		}

		s.Lock()
		s.connnected = true
		s.Unlock()
	}
	return nil
}

func GetDB() (*xorm.Engine, error) {
	i, _ := GetInstance()
	return i.engine, nil
}

func InitStore(ip string, port int, user string, password string, dbname string) error {
	instance,err := GetInstance()
	if err != nil {
		return err
	}
	return instance.Init(ip,port,user,password,dbname)
}
