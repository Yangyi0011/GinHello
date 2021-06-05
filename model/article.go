package model

import (
	"GinHello/initDB"
	"errors"
)

// Article 文章模型
type Article struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

// TableName 继承 TableName 接口，实现 gorm 对数据表名的指定
// 否则 gorm 默认会去操作 articles
func (article Article) TableName() string {
	return "article"
}

// Insert 插入文章
func (article *Article) Insert() (int, error) {
	create := initDB.DB.Create(&article)
	if create.Error != nil {
		return -1, create.Error
	}
	return article.Id, nil
}

// Update 更新文章
func (article *Article) Update() error {
	var a Article
	// 先查询一条记录保存在模型变量 a 中
	// 同：select * from article where (id = article.Id) limit 1
	where := initDB.DB.Where("id = ?", article.Id).Take(&a)
	if a.Id == 0 || where.Error != nil {
		return errors.New("对象不存在或已被删除：" + where.Error.Error())
	}

	// 更新整体
	// 同：update `article` set (`type` = article.Type), (`content` = article.Content)
	//		where (`article`.`id` = article.Id)
	save := initDB.DB.Save(article)
	if save.Error != nil {
		return save.Error
	}

	// 如果想更新 a 中的指定字段
	// 同 UPDATE `article` SET `type` = 'kotlin'  WHERE `article`.`id` = a.id
	//initDB.DB.Model(&a).Update("type", "kotlin")

	return nil
}

// FindAll 查询所以文章
func (article Article) FindAll() ([]Article, error) {
	var articles []Article
	find := initDB.DB.Find(&articles)
	if find.Error != nil {
		return nil, find.Error
	}
	return articles, nil
}

// FindById 通过 id 查找文章
func (article Article) FindById() (Article, error) {
	// 以 article.Id 作为 where 查询的过滤条件。该方法仅适用于主键为 int 类型时
	first := initDB.DB.First(&article, article.Id)
	return article, first.Error
}

// DeleteOne 通过 id 删除文章
func (article Article) DeleteOne() error {
	a, e := article.FindById()
	if e != nil || a.Id == 0 {
		return errors.New("对象不存在或已被删除")
	}
	// Delete 传的是值而不是指针，因为删除并不需要给对象赋值
	del := initDB.DB.Delete(article, article.Id)
	if del.Error != nil {
		return del.Error
	}
	return nil
}
