package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

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

func PrepareMysql() {
	sqlStr := "INSERT INTO user(id, name, age) VALUES(?,?,?);"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare err:", err)
	}

	rows, err := stmt.Exec(14, "我叫张大锤你叫什么大锤", 35)
	if err != nil {
		fmt.Println("insert err:", err)
	}
	fmt.Println("insert succ:", rows)

	rows, err = stmt.Exec(15, "石大锤", 25)
	if err != nil {
		fmt.Println("insert err:", err)
	}
	fmt.Println("insert succ:", rows)
}

func TransactionMysql() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("open transaction err: ", err)
	}

	// 事务开启成功
	sqlStr1 := "update user set age=10 where id = 10;"
	sqlStr2 := "update users set age=110 where id = 11;"

	res, err := tx.Exec(sqlStr1)
	if err != nil {
		fmt.Println("sql1 exec err: ", err)
		tx.Rollback()
		return
	}
    fmt.Println("sql1 exec succ: ", res)

	res, err = tx.Exec(sqlStr2)
	if err != nil {
		fmt.Println("sql2 exec err: ", err)
		tx.Rollback()
		return
	}
	fmt.Println("sql2 exec err: ", res)
    
	tx.Commit()

}

func main() {
	ConnMysql()
    // PrepareMysql()
	TransactionMysql()
}