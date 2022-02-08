package main

import (
	"fmt"
	"sync"

	// "gobase/logagent/kafaka"
	"gobase/logagent/config"
	"gobase/logagent/taillog"
	"time"

	"gopkg.in/ini.v1"
)

var (
    wg sync.WaitGroup
	appCfg = new(config.AppConfig)
)

func Run() {
	defer wg.Done()
    tailsLines := taillog.GetTailChan()
    for {
		select {
		case line, ok := <-tailsLines:
			fmt.Println("tails.Lines: ", line, ok)
			if !ok {
				fmt.Println("file close.")
				time.Sleep(time.Second)
				continue
			}
			fmt.Println("line: ", line.Text)
		default:
			time.Sleep(time.Second)
			fmt.Println("暂时没有数据写入.")
		}
	}
}

func main() {
	wg.Add(1)
    // 0. 加载配置文件
    err := ini.MapTo(appCfg, "./config/conf.ini")
	if err != nil {
		fmt.Println("ini config err: ", err)
		return
	}
	fmt.Println("ini succ:", appCfg, appCfg.KafakApp.Address, appCfg.KafakApp.Topic, appCfg.TailLogApp.FailPath)

	// 1. 初始化kafka连接
	// err := kafak.Init([]string{"127.0.0.1:9092"})
	// if err != nil {
	// 	fmt.Println("connect kafka err: ", err)
	// 	return
	// }

	// 2. 打开日志文件准备收集日志
	err = taillog.Init("./log/BackupService.log")
    if err != nil {
		fmt.Println("Init taillog err: ", err)
		return
	}

	go Run()
	wg.Wait()
}