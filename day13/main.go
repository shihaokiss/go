package main

import (
	"fmt"	
	"sync"
)


var (
    c = make(chan int, 10)
	wg sync.WaitGroup
	once sync.Once
)

func test() {
	defer wg.Done()
	c <- 10
	num := <- c
	fmt.Println(num)
	f := func() {
		close(c)
	}
	once.Do(f)
}

func test2() {
	for i:=0; i<10; i++{
		wg.Add(1)
		go test()
	}
	wg.Wait()
}

func main() {
	test2()
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var (
// 	wg      sync.WaitGroup
// 	rw_lock sync.RWMutex
// 	lock    sync.Mutex
// 	num = 0
// )

// func Read() {
//     defer wg.Done()
//     // lock.Lock()
// 	rw_lock.RLock()
// 	time.Sleep(time.Millisecond*1000)
// 	fmt.Println(num)
// 	// lock.Unlock()
// 	rw_lock.RUnlock()
// }

// func Write() {
// 	defer wg.Done()
// 	// lock.Lock()
// 	fmt.Println("i am write")
// 	rw_lock.Lock()
// 	num += 1
// 	time.Sleep(time.Millisecond*10)
// 	// lock.Unlock()
// 	rw_lock.Unlock()
// }

// func test() {
// 	start := time.Now()
//     for i:=0; i<10; i++ {
// 		go Write()
// 		wg.Add(1)
// 	}
//     time.Sleep(time.Second)
// 	for i:=0; i<100000; i++ {
// 		go Read()
// 		wg.Add(1)
// 	}

// 	wg.Wait()
// 	use_time := time.Now().Sub(start)
//     fmt.Println(use_time)
// }

// func main() {
// 	test()
// }
/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import (
// 	"fmt"
// 	"sync"
// )

// var x = 0
// var wg sync.WaitGroup
// var lock sync.Mutex
// func Add(count int) {
// 	defer wg.Done()
// 	for i := 0; i < count; i++ {
// 		// 加锁
// 		lock.Lock()
// 		x += 1
// 		lock.Unlock()
// 	}
// }

// func test() {
// 	wg.Add(2)
// 	go Add(1000000)
// 	go Add(1000000)
// 	wg.Wait()

// 	fmt.Println(x)
// }

// func main() {
// 	test()
// }