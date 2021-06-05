package model

import (
	"GinHello/initDB"
	"github.com/jinzhu/gorm"
)

// Comment 评论模型，建表时采用 orm 自动生成方式
type Comment struct {
	// 这里 ID 作为主键要全部大写
	// 如果在建表时想同时建立索引，则可以为指定字段添加 tag。如：`gorm:"index:idx_id"`
	//ID      int

	// 继承 gorm.Model，引入必要的字段，如：ID、创建时间、修改时间、删除时间
	gorm.Model
	Content string
}

// 项目运行后会自动建表
func init() {
	table := initDB.DB.HasTable(Comment{})
	// 表不存在时才创建
	if !table {
		initDB.DB.CreateTable(Comment{})
	}
}