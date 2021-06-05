package model

import (
	"GinHello/initDB"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	// 继承 gorm.Model 以生存必要字段
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Roles    []Role `json:"role" gorm:"many2many:roles"`
}

func init() {
	// 表不存在的时候创建表
	// 不继承 TableName() string 方法来指定表名的话，默认创建的是 users
	if !initDB.DB.HasTable(User{}) {
		initDB.DB.CreateTable(User{})
	}
}

// QueryByUsername 通过 username 查询用户信息
func (user User) QueryByUsername() User {
	initDB.DB.First(&user, user.Username)
	return user
}

// Insert 添加用户
func (user User) Insert() bool {
	user.CreatedAt = time.Now()
	create := initDB.DB.Create(&user)
	if create.Error != nil {
		return false
	}
	return true
}
