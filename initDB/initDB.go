package initDB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 配置信息常量
const (
	driverName = "mysql"
	username   = "root"
	password   = "toor"
	host       = "localhost"
	port       = "3306"
	dataBase   = "ginhello"

	// 最大连接数
	maxOpenConns = 10
	// 最大空闲连接数
	maxIdleConns = 1
)

var (
	// DataBase 数据库连接变量
	DataBase *sql.DB
)

// 初始化数据源信息
func init() {
	var err error
	dataSourceName := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dataBase
	DataBase, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Panicln("DataBase connection error -> ", err.Error())
	}
	DataBase.SetMaxOpenConns(maxOpenConns)
	DataBase.SetMaxIdleConns(maxIdleConns)
}
