package main

import "fmt"

func test() {
	slice := make([]int, 0)
	slice = append(slice, 1, 2, 3)
	var I interface{} = slice
	fmt.Println(I.([]int))
	res, ok := I.([]float32)
	fmt.Println(res, ok)
	if ok {
		fmt.Println(res)
		fmt.Println(ok)
	}
}

func main() {
	test()
}

// import (
// 	"fmt"
// )

// type Num int

// func (x Num) Equal(i int) bool {
// 	return x == i
// }

// func (x *Num) Multiple(i int) {
// 	*x = *x * i
// }

// func test() {
//     var num Num = 6
// 	var tt1 T1.MyInterface1 = num
// 	var tt2 T2.MyInterface2 = num
// 	result := tt1.Equal(6)
// 	fmt.Println(result)

// 	result = tt2.Multiple(8)
// 	fmt.Println(result, num)
// }

// type Num int

// func (x Num) Equal(i Num) bool {
// 	return x == i
// }

// import "fmt"

// type Num int

// func (x Num) Equal(i Num) bool {
// 	return x == i
// }

// func (x *Num) Multiple(i Num) {
// 	*x = *x * i
// }

// type NumI interface {
// 	Equal(i Num) bool
// 	Multiple(x Num)
// }

// func test() {
// 	var x Num = 8
// 	var y NumI = &x
// 	data := y.Equal(8)
//     fmt.Println(data)
// 	y.Multiple(6)
// 	fmt.Println(x)
// }

// func main() {
// 	test()
// }

// import (
// 	alias2 "fmt"
// )

// func test() {
//     type mt interface {
// 		Run()
// 		Stop()
// 	}

// 	type Bus struct {
// 		mt
// 	}

// }

// func main() {
//     test()
// }

// func test() {
// var name string = "zhangsan"
// var age int = 10
// var (
// 	name string = "zhangsan"
// 	age int = 10
// )
// name, age := "zhangsan", 10
// alias2.Println(name, age)
// const changliang int = 10
// alias2.Println(changliang)
// const changliang2 = os.Getenv("url")
// alias2.Println(changliang2)

// num := 1
// for i:=1; i<5; i++ {
// 	num *= i
// }
// alias2.Println(num)

// num, i := 0, 0

// for {
// 	i++
// 	num += i
// 	if i > 50 {
// 		break
// 	}
// }
// alias2.Println(num)

// ss := "Let't go."
// for i, c := range ss {
// 	alias2.Println(i, c)
// }

// ll := []int {1 , 2, 3, 4, 5}
// for key, value := range ll {
// 	alias2.Println(key, value)
// }
// m := map[string]int {
// 	"aaa": 10,
// 	"bbb": 20,
// }
// for key, value := range m {
// 	alias2.Println(key, value)
// }

// c := make(chan int)
// go func() {
// 	c <- 10
// 	c <- 20
// 	c <- 30
// 	close(c)
// }()
// for v := range c {
// 	alias2.Println(v)
// }

// a := "love3"
// switch a {
// case "love", "love2":
// 	alias2.Println("i love ...")
// default:
// 	alias2.Println("default..")
// }

// num := 30
// switch {
// case num < 10:
// 	alias2.Println("num <10")
// case num < 20:
// 	alias2.Println("10<= num <20")
// default:
// 	alias2.Println("err num")
// }

// OutLoop:
//     for i:=0; i<2; i++ {
// 		for j:=0; j<3; j++ {
//             switch j {
// 			case 1:
// 				alias2.Println(i, j)
// 				continue OutLoop
// 			}
// 		}
// 	}

// ss := "中国"
// l1 := len(ss)    //字节数
// l2 := utf8.RuneCountInString(ss)    //字符个数
// alias2.Println(l1, l2)

// sss := "我爱中国"
// // chars := []rune(sss)
// // alias2.Println(len(chars))
// for i, char := range sss {
// 	alias2.Println(i, string(char))
// }

// s1 := "i love 中国"
// by1 := []byte(s1)
// alias2.Println(len(by1))
// by1[0] = 'y'
// alias2.Println(string(by1))

// s1 := "i love 中国"
// by1 := []rune(s1)
// alias2.Println(len(by1))
// by1[7] = '世'
// by1[8] = '界'
// alias2.Println(string(by1))

// a, b := 1, 2
// pa, pb := &a, &b
// alias2.Println(pa, pb)
// pa, pb = pb, pa
// alias2.Println(pa, pb)
// alias2.Println(*pa, *pb)
// alias2.Println(a, b)

// names := [5]int {1,2,3,4,5}
// alias2.Println(names)

// type Book struct {
// 	name string
// 	author string
// }

// b1 := Book{name: "go web", author: "shihao"}
// alias2.Println(b1)

// 声名切片
// var sliceStr []string
// alias2.Println(sliceStr, len(sliceStr), sliceStr == nil)
// sliceStr := make([]string, 5, 1000)
// alias2.Println(len(sliceStr))

// map
// myMap := make(map[string]string, 10)
// myMap["aaa"] = "111"
// myMap["bbb"] = "222"
// alias2.Println(myMap)
// }

// func main() {
// 	test()
// }

// import (
// 	"fmt"
// 	"time"
// )

// func test_chan2(c chan int){
// 	fmt.Println("hello1")
//     time.Sleep(time.Second*1)
// 	fmt.Println("hello2")
// 	data := <- c
//     fmt.Println("test_chan2", data)

// 	close(c)

// 	data = <- c
//     fmt.Println("test_chan2", data)
// 	data = <- c
//     fmt.Println("test_chan2", data)
// 	data = <- c
//     fmt.Println("test_chan2", data)
// 	data = <- c
//     fmt.Println("test_chan2", data)

// 	close(c)
// }

// // 无缓冲通道也成同步通道
// func test_chan(){
// 	my_chan := make(chan int)
// 	go test_chan2(my_chan)
// 	my_chan <- 10
//     fmt.Println("test_chan send end.")
// 	// close(my_chan)
// 	// my_chan <- 100
// }

// func main(){
// 	test_chan()
// }

// 测试切片
// func test_slice(){
// 	a := []int{}
// 	fmt.Println(a, len(a), cap(a))

// 	b := []int{1,2,3}
// 	fmt.Println(b, len(b), cap(b))

// 	c := b[0:2]
// 	fmt.Println(c, len(c), cap(c))
// }

// func main(){
// 	test_slice()
// }

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
