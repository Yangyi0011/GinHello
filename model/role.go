package model

import (
	"GinHello/initDB"
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name string `json:"name"`
}

func init() {
	// 表不存在的时候创建表
	// 不继承 TableName() string 方法来指定表名的话，默认创建的是 roles
	if !initDB.DB.HasTable(Role{}) {
		initDB.DB.CreateTable(Role{})
	}
}
