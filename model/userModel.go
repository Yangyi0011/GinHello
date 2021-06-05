package model

import (
	"GinHello/initDB"
	"database/sql"
	"log"
)

// UserModel 用户模型
type UserModel struct {
	Id int  `form:"id"`
	// 通过 form:"email" 来对表单中的 email 输入数据进行绑定
	// binding 标签是用来进行数据校验的，binding:"email" 表示该数据必须是邮箱格式
	// Gin 对于数据校验使用的是 validator.v8 库，该库提供多种校验方法。
	Email string `form:"email" binding:"email"`
	Password string `form:"password"`
	// binding:"eqfield=Password" 表示 RePassword 的值必须与 Password 一致
	//RePassword string `form:"rePassword" binding:"eqfield=Password"`

	// 为什么是 sql.NullString ？因为我们数据库中该字段初始时为 null ，
	// 而 string 类型是不可以接收 null 类型的，所以我们只能采用 NullString
	// 来对 null 字符串进行处理。
	Avatar sql.NullString `form:"avatar"`
}

// Save 保存用户数据，返回该用户数据的 id
func (user *UserModel) Save() int64 {
	sqlStr := "Insert Into ginhello.user(email, password) values(?,?);"
	result, e := initDB.DataBase.Exec(sqlStr, user.Email, user.Password)
	if e != nil {
		log.Panicln("user insert error ->", e.Error())
	}
	// 获取刚刚插入数据的 id
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("user insert id error ->", err.Error())
	}
	return id
}

// QueryByEmail 通过 email 查询用户信息
func (user *UserModel) QueryByEmail() UserModel {
	u := UserModel{}
	sqlStr := "select * from user where email = ?"
	row := initDB.DataBase.QueryRow(sqlStr, user.Email)
	err := row.Scan(&u.Id, &u.Email, &u.Password, &u.Avatar)
	if err != nil {
		log.Panicln(err)
	}
	return u
}

// QueryById 通过 id 查询用户信息
func (user *UserModel) QueryById(id int) (UserModel, error) {
	u := UserModel{}
	sqlStr := "select * from user where id = ?"
	row := initDB.DataBase.QueryRow(sqlStr, id)
	e := row.Scan(&u.Id, &u.Email, &u.Password, &u.Avatar)
	if e != nil {
		log.Panicln(e)
	}
	return u, e
}

// Update 更新用户信息
func (user *UserModel) Update(id int) error {
	sqlStr := "update user set password=?,avatar=?  where id=? "
	// sql 预处理
	var stmt, e = initDB.DataBase.Prepare(sqlStr)
	if e != nil {
		log.Panicln("发生了错误", e.Error())
	}
	// 执行 sql
	_, e = stmt.Exec(user.Password, user.Avatar.String, id)
	if e != nil {
		log.Panicln("错误 e", e.Error())
	}
	return e
}