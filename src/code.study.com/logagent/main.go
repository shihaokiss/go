package main

import (
	"fmt"
	"code.study.com/logagent/taillog"
)


func main() {
	// 1. 初始化Kafka
	fmt.Println("Init Kafka success.")
	// 2. 打开日志文件准备收集日志
	// err := taillog.Init("./my.log")
	// if err != nil {
	// 	fmt.Println("Init Log err:", err)
	// 	return
	// }
	// taillog.ShowInfo()
	fmt.Println("Init Log success.")
}
