package main

import "fmt"

func testSlice() {
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)

	var slice2 []int
	slice2 = make([]int, 3)
	fmt.Println(slice2)

	var slice3 []int = make([]int, 3)
	fmt.Println(slice3)

	slice4 := make([]int, 3)
	fmt.Println(slice4)

	var slice5 []int
	if slice5 == nil {
		fmt.Println("empty slice")
	} else {
		fmt.Println("not empty slice", slice5)
	}
}

func main() {
	testSlice()
}
