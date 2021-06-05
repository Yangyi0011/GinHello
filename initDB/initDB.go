package initDB

import (
	"GinHello/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var (
	// DB 数据源连接变量
	DB *gorm.DB
)

// 初始化数据源信息
func init() {
	var err error
	dataSourceName := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.DataBase
	DB, err = gorm.Open(config.DriverName, dataSourceName)
	if err != nil {
		log.Panicln("Datasource connection error -> ", err.Error())
	}
}
