package main

import (
	"fmt"
	"strconv"
	"sync"
)

func test(num int) {
	m1 := sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i:=0; i<num; i++ {
		go func(n int) {
            defer wg.Done()
			s := strconv.Itoa(n)
			m1.Store(s, n)
			res, ok := m1.Load(s)
			fmt.Println("key=", s, "value=", res, ok)
		}(i)
	}
    wg.Wait()
	fmt.Print("end...")
}

func main() {
	test(20)
}
