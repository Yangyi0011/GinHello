package initDB

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
	// DB 数据源连接变量
	DB *gorm.DB
)

// 初始化数据源信息
func init() {
	var err error
	dataSourceName := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dataBase
	DB, err = gorm.Open(driverName, dataSourceName)
	if err != nil {
		log.Panicln("Datasource connection error -> ", err.Error())
	}
}
