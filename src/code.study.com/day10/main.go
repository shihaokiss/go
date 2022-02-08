package main

import (
	"fmt"
	"strings"
)

// 闭包
// 使用场景：将有两个参数的函数fun3当作参数，传给没有参数的函数func1
func func1(f func()) {
    f()
}

func func2(f func(x, y int), a, b int) func() {
	tmp := func() {
        f(a, b)
	}
	return tmp
}

func func3(a, b int) {
    fmt.Println("func3 ", a+b)
}

// 在一个函数中使用外部作用域的变量
func func11(x int) func(int) int {
	return func(y int) int {
        fmt.Println("func11", x+y)
		x += y
		return x
	}
}

// 给文件名后面加后缀
func func22(houzhui string) func(string) string{
	return func(name string) string {
        if !strings.HasSuffix(name, houzhui) {
			return name + houzhui
		} else {
			return name
		}
	}
}

// 给一个base 数字，做数字运算
func calc(base int) (func(int) int, func(int) int){
	add := func(num int) int{
		base += num
		return base
	}

	sub := func(num int) int {
		base -= num
		return base
	}

	return add, sub
}

func main() {
	fmt.Println("I am main")
	// func1(func2(func3, 1, 2))
	// f1 := func11(10)
	// fmt.Println(f1(20))
	// fmt.Println(f1(20))

	// add_txt := func22(".txt")
	// add_jpg := func22(".jpg")
	// fmt.Println(add_txt("aaaa"))
	// fmt.Println(add_jpg("aaaa"))

	add, sub := calc(10)
	fmt.Println(add(10), sub(2))
}