package main

import (
	"fmt"
	"runtime"
	"sync"
	// "time"
)

// func show(i int) {
// 	fmt.Println("show", i)
// }

// func test() {
// 	for i := 0; i < 100; i++ {
//         go func(num int) {
//             fmt.Println("show", num)
// 		}(i)
// 	}
// }

// func main() {
// 	test()
// 	fmt.Println("main")
// 	time.Sleep(time.Second)
// }

// wait 等待
// var wg sync.WaitGroup
// func do_handle(i int) {
// 	defer wg.Done()
//     fmt.Println("do_handle", i)
// }

// func test() {
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go do_handle(i)
// 	}
// }

// func main() {
// 	test()
// 	wg.Wait()
// 	fmt.Println("main")
// }

// 设置占用核心数，GOMAXPROCES
var wg sync.WaitGroup

func A() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("A", i)
	}
}

func B() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("B", i)
	}
}

func test() {
	runtime.GOMAXPROCS(256)
	wg.Add(2)
	go A()
	go B()
	wg.Wait()
}

func main() {
	test()
}



















