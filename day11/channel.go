package main

import (
	"fmt"	
	// "sync"
)


// 通道的使用
// var wg sync.WaitGroup

// func noBufChannel() {
// 	var c1 chan int            //声名
// 	c1 = make(chan int)        //初始化
    
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		num := <-c1
// 		fmt.Println("后台程序成功接收，", num)
// 	}()
    
// 	fmt.Println("向通道中发送信号。")
// 	c1 <- 10
// 	// c2 := make(chan int, 20)   //带缓冲区的通道
// 	fmt.Println(c1)
// }

// func bufChannal() {
// 	var c1 chan int
// 	c1 = make(chan int, 2)
// 	c1 <- 10
// 	fmt.Println("10 in")
// 	c1 <- 20
// 	fmt.Println("20 in")

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 3; i++ {
// 			num := <-c1
// 			fmt.Println("后台程序成功接收，", num)
// 		}
// 	}()
// 	c1 <- 30


// 	println("i will out.")
// }


// func main() {
// 	bufChannal()
// }


// channel 练习
// 1. 向通道 c1 中写 100 个数
// 2. 后台 goroutine 从 c1 中获取
// 3. 将获取到的数字进行平方放到 c2 中
// 4. 主程序从 c2 中获取对应的数字
// var wg sync.WaitGroup

// func test() {
// 	c1 := make(chan int, 100)
// 	c2 := make(chan int, 100)
    
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 10; i++ {
// 			num := <- c1
// 			num2 := num*num
// 			fmt.Println("向管道c2放，", num2)
// 			c2 <- num2
// 		}
// 	}()
    
// 	go func () {
// 		defer wg.Done()
// 		for i := 0; i < 10; i++ {
// 			fmt.Println("向管道c1放，", i)
// 			c1 <- i
// 		}
// 	}()

// 	for i := 0; i < 10; i++ {
// 		num := <- c2
// 	    fmt.Println("从c2中取数据，", num)
// 	}

// 	fmt.Println("test exit.")
// }

// func main() {
// 	test()
// }


func test() {
	var c1 chan int
	c1 = make(chan int, 10)

	c1 <- 10
	c1 <- 20
	c1 <- 30

	num, ok := <-c1
	fmt.Println(num, ok)

	num, ok = <-c1
	fmt.Println(num, ok)

	num, ok = <-c1
	fmt.Println(num, ok)
    close(c1)
	num, ok = <-c1
	fmt.Println(num, ok)
}

func main() {
	test()
}
















