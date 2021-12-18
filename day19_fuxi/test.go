package main

import "fmt"


func test_slice(){
	a := []int{}
	fmt.Println(a, len(a), cap(a))
    
	b := []int{1,2,3}
	fmt.Println(b, len(b), cap(b))

	c := b[0:2]
	fmt.Println(c, len(c), cap(c))
}

func main(){
	test_slice()
}

// 一个简单的求和函数
// func sum(l [3]int) int{
// 	sum := 0
// 	for i, v := range l{
//         fmt.Println(i, v)
// 		sum += v
// 	}
// 	return sum
// }

// func main() {
// 	l := [3]int{1, 2, 3}
// 	sum := sum(l)
// 	fmt.Println(sum)
// }

// 数字表达的另一种方式
// func test() {
// 	v := 123_456
// 	fmt.Printf("num = %d", v)
// }

// func main() {
// 	test()
// }