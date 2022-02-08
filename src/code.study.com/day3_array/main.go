package main

import "fmt"

func arrayTest(p_array [5]int) {
	// 值传递
	for i := 0; i < len(p_array); i++ {
		fmt.Println(i, p_array[i])
	}
	fmt.Println("++++++++++++++")
	for i, value := range p_array {
		fmt.Println(i, value)
	}
	p_array[0] = 100
	fmt.Println(p_array)
}

func arraySlice(p_array []int) {
	for i, value := range p_array {
		fmt.Println(i, value)
	}
	p_array[0] = 100
}

func test1() {
	array2 := [5]int{1, 2, 3, 4, 5}
	arrayTest(array2)
	fmt.Println(array2)
}

func test2() {
	array1 := []int{1, 2, 3, 4, 5} //slice 动态数组
	fmt.Println(array1)
	arraySlice(array1)
	fmt.Println(array1)
}
func main() {
	// test1()
	test2()
}
