///////////////////////////////////////////////////测试通道range/////////////////////////////////////////////////////
package main

import (
	"fmt"
	"time")



func test() {
	chan1 := make(chan int, 10)

	go func(c chan int) {
        for i:=0; i< 10; i++{
			c <- 10
		}
        time.Sleep(time.Second*1000)
        for i:=0; i< 10; i++{
			c <- i
		}
	}(chan1)


	// close(chan1)
	for {
		i, ok := <- chan1
		fmt.Printf("i=%d\n", i)
		// 当管道关闭并且管道中为空的时候 ok=false
		if !ok {
			break
		}
	}
	
	for i := range(chan1) {
		fmt.Printf("i=%d\n", i)
	}

	time.Sleep(time.Millisecond*2000)
}

func main() {
	test()
}



///////////////////////////////////////////////////context 练习：子goroutine退出时让它的子goroutine退出/////////////////////////////////////////////////////
// package main

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var wg sync.WaitGroup

// func worker2(ctx context.Context) <-chan int {
// 	defer fmt.Println("worker2 exit.")
// 	chan1 := make(chan int, 10)
// 	chan2 := make(chan int, 10)
// 	num := 0
// 	go func() {
// 		defer wg.Done()
// 		defer fmt.Println("worker2 func exit.")
// 		for {
//             select {
// 			case chan1<-num:
// 				fmt.Println("worker2 func chan1 put:", num)
// 				num = num + 1
// 				time.Sleep(time.Second)
// 			case chan2<-num:
// 				fmt.Println("worker2 func chan2 put:", num)
// 				time.Sleep(time.Second)
// 			case <-ctx.Done():
// 				return
// 			}
// 		}
// 	}()
// 	return chan1
// }

// func worker1() {
// 	wg.Add(1)
// 	ctx, cancle := context.WithCancel(context.Background())
// 	defer fmt.Println("worker1 exit.")
	
// 	for n := range worker2(ctx) {
// 		fmt.Println("worker1 get:", n)
//         if n == 5 {
// 			break
// 		}
// 	}
// 	cancle()
// 	wg.Wait()
// }

// func main() {
// 	worker1()
// }




///////////////////////////////////////////////////context 练习：通知子goroutine退出/////////////////////////////////////////////////////
// package main

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var wg sync.WaitGroup

// func worker2(ctx context.Context) {
// 	defer wg.Done()
// 	defer fmt.Println("worker2 exit2.")
// 	defer fmt.Println("worker2 exit1.")
// 	for {
// 		fmt.Println("worker2")
// 		time.Sleep(time.Millisecond * 500)
// 		select {
// 		case <-ctx.Done():
// 			return
// 		default:
// 		}
// 	}
// }

// func worker1(ctx context.Context) {
// 	defer wg.Done()
// 	defer fmt.Println("worker1 exit2.")
// 	defer fmt.Println("worker1 exit1.")
// 	go worker2(ctx)
// 	for {
// 		fmt.Println("worker1")
// 		time.Sleep(time.Millisecond * 500)
// 		select {
// 		case <-ctx.Done():
// 			return
// 		default:
// 		}
// 	}
// }

// func main() {
// 	wg.Add(2)
// 	ctx, cancle := context.WithCancel(context.Background())
// 	go worker1(ctx)
// 	time.Sleep(time.Second * 2)
// 	// 通知父 context 下面所有的子 context 超时
// 	cancle()
// 	wg.Wait()
// }

////////////////////////////////////////////////////////////////////////////////////////////////////////
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var wg sync.WaitGroup
// var exitInfo = make(chan bool, 10)

// func f() {
// 	defer wg.Done()
// 	defer fmt.Println("f() exit.")
// 	for {
// 		fmt.Println("test")
// 		time.Sleep(time.Millisecond*500)
// 		select {
// 		case <-exitInfo:
// 			return
// 		default:
// 			fmt.Println("select default;")
// 		}
// 	}
// }

// func main() {
// 	wg.Add(1)

// 	go f()
// 	time.Sleep(time.Second*5)
// 	exitInfo <- true

//     wg.Wait()
// }
