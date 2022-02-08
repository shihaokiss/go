package main

import (
	"fmt"
	// "runtime"
	"time"
)

// 死循环函数
func RunShowNum() {
	i := 0
	for {
		i++
		fmt.Println("RunShowNum num=", i)
		time.Sleep(1 * time.Second)
	}
}


// 匿名函数，中嵌套推出go程
func test() {
	go func(a int, b int) {
        defer fmt.Println("func A End.")
		func() {
			defer fmt.Println("func B End.")
			//推出当前 goroutine
			// runtime.Goexit()
			fmt.Println("func B")
		}()
		fmt.Println("func A ", a, b)
	}(10, 20)
}


func main() {
	// 运行go程
	// go RunShowNum()
	test()

	time.Sleep(1 * time.Second)
}