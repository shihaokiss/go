package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("A::func::defer")

		func() {
			defer fmt.Println("B::func::defer")
			fmt.Println("B::func")
			return
		}()

		fmt.Println("A::func")
		return
	}()

	go func(a int, b int) bool {
		fmt.Println("china::a::b", a, b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}
