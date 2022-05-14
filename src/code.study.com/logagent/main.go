package main

import (
	"fmt"
	"time"
	"code.study.com/logagent/conf"
	"code.study.com/logagent/kafka"
	"code.study.com/logagent/taillog"
	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

func run() {
	for {
		select {
		case line := <-taillog.ReadChan():
			fmt.Println("taillog get info :", line.Text)
			kafka.SendToKafka(cfg.Kafka.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
    // 0. 加载配置文件
	err := ini.MapTo(cfg, "./conf/conf.ini")
	if err != nil {
        fmt.Println("Init Kafka faile:", err)
		return
	}
    fmt.Println("Init Conf success.")

	// 1. 初始化Kafka
	err = kafka.Init([]string{cfg.Kafka.Address})
	if err != nil {
		fmt.Println("Init Kafka faile:", err)
		return
	}
	fmt.Println("Init Kafka success.")

	// 2. 打开日志文件准备收集日志
	err = taillog.Init(cfg.TailLog.FileName)
	if err != nil {
		fmt.Println("Init taillog failed: ", err)
		return
	}
	fmt.Println("Init Log success.")

	run()
}
