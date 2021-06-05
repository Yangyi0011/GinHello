package model

import (
	"GinHello/initDB"
	"errors"
	"log"
)

// Article 文章模型
type Article struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

// Insert 插入文章
func (article *Article) Insert() (int, error) {
	sqlStr := "insert into article(type, content) values(?,?);"
	result, err := initDB.DataBase.Exec(sqlStr, article.Type, article.Content)
	if err != nil {
		log.Println("文章添加失败", err.Error())
		return -1, err
	}
	id, e := result.LastInsertId()
	return int(id), e
}

// Update 更新文章
func (article *Article) Update() error {
	sqlStr := "update article set type = ?, content = ? where id = ?"
	// 执行 sql
	_, e := initDB.DataBase.Exec(sqlStr, article.Type, article.Content, article.Id)
	if e != nil {
		log.Println("更新失败：", e.Error())
		return e
	}
	return nil
}

// FindAll 查询所以文章
func (article Article) FindAll() ([]Article, error) {
	sqlStr := "select * from article;"
	rows, e := initDB.DataBase.Query(sqlStr)
	if e != nil {
		log.Println("查询数据失败")
		return nil, e
	}

	var articles []Article
	for rows.Next() {
		var a Article
		if e := rows.Scan(&a.Id, &a.Type, &a.Content); e != nil {
			log.Println("封装数据失败")
			return nil, e
		}
		articles = append(articles, a)
	}
	return articles, nil
}

// FindById 通过 id 查找文章
func (article Article) FindById() (Article, error) {
	sqlStr := "select * from article where id=?;"
	row := initDB.DataBase.QueryRow(sqlStr, article.Id)
	if e := row.Scan(&article.Id, &article.Type, &article.Content); e != nil {
		log.Println("封装数据失败", e.Error())
		return article, e
	}
	return article, nil
}

// DeleteOne 通过 id 删除文章
func (article Article) DeleteOne() error {
	a, e := article.FindById()
	if e != nil || a.Id == 0 {
		return errors.New("对象不存在或已被删除")
	}
	sqlStr := "delete from article where id = ?"
	if _, e := initDB.DataBase.Exec(sqlStr, article.Id); e != nil {
		log.Println("删除失败")
		return e
	}
	return nil
}
