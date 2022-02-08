package main

// import (
// 	// "fmt"
// 	// "time"
// )
// // 从通道中取值
// func test() {
// 	c := make(chan int, 10)
// 	for i := 0; i < 10; i++ {
// 		c <- i
// 	}
//     close(c)
//     c <- 10
// 	// for i := range(c) {
// 	//     fmt.Println("test: ", i)
// 	// }
//
// 	// for {
// 	// 	num, ok := <-c
// 	// 	if ok {
// 	// 		fmt.Println("test_num_ok ", num, ok)
// 	// 		time.Sleep(time.Second)
// 	// 	} else {
// 	// 		fmt.Println("test_num_ok ", num, ok)
// 	// 		time.Sleep(time.Second)
// 	// 	}
// 	// }
// }
//
// func main() {
// 	test()
// }

// import (
// 	"fmt"
// 	"time"
// )
//
// // select 多路复用
// func test() {
//     //c := make(chan int)    错误写法，无缓冲通道会阻塞，直到有人去读
//     c := make(chan int, 1)    // 正确写法, 读写不会被阻塞
//     for i:=0; i<10; i++ {
//         select {
//         case c <- i:
//             fmt.Println("input ", i)
//         case num := <-c:
//             fmt.Println("get num ", num)
//         }
//     }
// }
//
// func main() {
//     test()
// }

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
// goroutine pool
func worker(jobs <-chan int, results chan<- int, index int) {
    defer wg.Done()
    for {
        job, is_close := <-jobs
        if is_close {
            fmt.Printf("worker: %d, start job: %d\n", index, job)
            time.Sleep(time.Second * 1)
            // fmt.Printf("worker: %d, end job: %d\n", index, job)
            results <- job * 2
        } else {
            fmt.Printf("我要关闭了index=%d, is_close=%t\n", index, is_close)
            time.Sleep(time.Second * 1)
            break
        }
    }
}

func test() {
    jobs := make(chan int, 10)
    results := make(chan int, 10)

    for i:=0; i<10; i++ {
        jobs <- i
    }

    // goroutine pool size = 3
    wg.Add(3)
    for j:=0; j<3; j++ {
        go worker(jobs, results, j)
    }

    close(jobs)

    wg.Wait()    //测试说明worker工作完成
    close(results)

    for num := range(results) {
        fmt.Println(num)
    }
    time.Sleep(time.Second * 5)
    fmt.Println("end..")
}

func main() {
    test()
}

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )
// 
// var wg sync.WaitGroup
// 
// func channel_test1(c1 chan<- int) {
// 	fmt.Println("channel_test1::c1 len=", len(c1))
// }
// 
// func main() {
//     wg.Add(1)
// 	c1 := make(chan int, 10)
// 	for i := 0; i < 10; i++ {
// 		c1 <- i
// 	}
//     fmt.Println("c1 len=", len(c1))
// 	// go channel_test1(c1)
//     go func(c1 chan int) {
//         fmt.Println("xx")
//         fmt.Println("func::c1 len=", len(c1))
//     }(c1)
// 
//     wg.Done()
//     time.Sleep(time.Second * 2)
// }
