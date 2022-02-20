package main

import (
	"fmt"
	"time"

	"code.study.com/logagent/kafka"
	"code.study.com/logagent/taillog"
)

func run() {
	for {
		select {
		case line := <-taillog.ReadChan():
			fmt.Println("taillog get info :", line.Text)
			kafka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 1. 初始化Kafka
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Println("Init Kafka faile:", err)
		return
	}
	fmt.Println("Init Kafka success.")

	// 2. 打开日志文件准备收集日志
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Println("Init taillog failed: ", err)
		return
	}
	fmt.Println("Init Log success.")

	run()
}
