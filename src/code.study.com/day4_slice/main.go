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

func sliceAppend() {
	var numbers = make([]int, 3, 5)
	fmt.Println(len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 1)
	fmt.Println(len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 1)
	fmt.Println(len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 1)
	fmt.Println(len(numbers), cap(numbers), numbers)

	var numbers2 = make([]int, 3)
	numbers2 = append(numbers2, 1)
	fmt.Println(len(numbers2), cap(numbers2), numbers2)
}

func sliceSplit() {
	s1 := []int{1, 2, 3}
	s2 := s1[0:2]
	fmt.Println(len(s2), cap(s2), s2)
	s2[0] = 100
	fmt.Println(s1, s2)

	s3 := make([]int, 4)
	copy(s3, s1)
	fmt.Println(len(s3), cap(s3), s3)
	s3[0] = 10000
	fmt.Println(s1, s3)
}

func main() {
	// testSlice()
	// sliceAppend()
	sliceSplit()
}
