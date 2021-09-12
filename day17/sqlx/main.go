package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

func ConnMysql() (err error) {
    // 用户名:密码@tcp(ip:port)/数据库名
	dsn := "root:@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True"
	// 校验 dsn 格式
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("db connect err: ", err)
	}
	db.SetMaxOpenConns(10)    // 设置连接池最大持有连接数
	db.SetMaxIdleConns(5)     // 设置连接池最大空闲连接数

	fmt.Println("sqlx 连接数据库成功。")
	return
}

type user struct {
	ID int
	Name string
	Age int
}

func GetUsers() {
	sql1 := "select * from user;"
	users := make([]user, 0, 10)

	err := db.Select(&users, sql1)
	if err != nil {
		fmt.Println("select err: ", err)
	}
	fmt.Println(users)
}

func main() {
	ConnMysql()
	GetUsers()
}