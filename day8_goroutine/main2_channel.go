package main

import (
	"fmt"
	"time"
)

func channel_no_buf() {
	fmt.Println("main start...")
	// 创建一个无缓冲int管道
	c := make(chan int)

	go func() {
		defer fmt.Println("func B end...")
		fmt.Println("func B running...")
		// 向管道中写 666
		c <- 666
	}()

	time.Sleep(1 * time.Second)
	num, ok := <-c
	fmt.Println("main get num=", num, ok)
}

func channel_buf() {
	fmt.Println("main start...")
	//创建一个容量为3的管道
	c := make(chan int, 3)

	go func() {
		defer fmt.Println("func B end...")
		fmt.Println("func B running...")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("func B c len=", len(c), " cap=", cap(c))
		}
	}()

	time.Sleep(1 * time.Second)

	for i := 0; i < 4; i++ {
		num, ok := <-c
		fmt.Println("main num=", num, ok)
	}
	fmt.Println("main end...")
	time.Sleep(1 * time.Second)
}

func fabonacii(c, quit chan int) {
	defer fmt.Println("func B end...")
	a, b := 0, 1
	for {
		// 管道控制协调 goroutine 的调度 ※※※※※
		select {
		case c <- a:
			// fmt.Println("func B ", a, b)
			tmp := a + b
			a = b
			b = tmp
			// fmt.Println("func B end")
			// time.Sleep(1 * time.Second)
		case <-quit:
			return
		}
	}
}
func channel_fibonacii() {
	fmt.Println("main start...")

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			num, ok := <-c
			fmt.Println("main ", num, ok)
		}
		quit <- 0
	}()

	fabonacii(c, quit)
	fmt.Println("main end...")
}

func main() {
	channel_fibonacii()
}
