package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Println("newTask::num::", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go newTask()
	fmt.Println("main::exit")
	// i := 0
	// for {
	// 	i++
	// 	fmt.Println("main::num", i)
	// 	time.Sleep(1 * time.Second)
	// }
}
