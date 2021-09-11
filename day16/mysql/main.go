package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type User struct {
	id string
	name string
	age int
}

func ConnMysql() {
    // 用户名:密码@tcp(ip:port)/数据库名
	dsn := "root:@tcp(127.0.0.1:3306)/godb"
	// 校验 dsn 格式
	db, err = sql.Open("mysql", dsn)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("dsn check err: ", err)
	}
	// 连接数据库
    err = db.Ping()
	if err != nil {
		fmt.Println("conn db err: ", err)
	}
	fmt.Println("conn succ")

	db.SetMaxOpenConns(10)    // 设置连接池最大持有连接数
	db.SetMaxIdleConns(5)     // 设置连接池最大空闲连接数
}

func QueryOne() {
	var user User
	sqlStr := "select * from user where id = ?"
	db.QueryRow(sqlStr, 1).Scan(&user.id, &user.name, &user.age)
	fmt.Printf("user=%#v\n", user)
}

func queryMutilRows() {
	sqlStr := "select * from user where id > 0;"
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("queryMutilRows err: ", err)
	}
    // 必须手动归还连接到连接池
	defer rows.Close()
	for rows.Next() {
		var user User
		err = rows.Scan(&user.id, &user.name, &user.age)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("user=%#v\n", user)
	}
}

func main() {
	ConnMysql()
	// QueryOne()
	queryMutilRows()
}