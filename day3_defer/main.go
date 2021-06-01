package main

import "fmt"

func func1() {
	fmt.Println("func1()...")
}

func func2() int {
	fmt.Println("func2()...")
	return 0
}

func func3() {
	fmt.Println("func3()...")
}

func deferFunc() {
	defer func1()
	defer func2()
	defer func3()
}

func deferAndReturn() int {
	defer func1()
	return func2()
}

func main() {
	// deferFunc()    defer 执行顺序
	deferAndReturn() // defer 和 return 的执行顺序。defer 类似于 finally、析构函数。在return之后
}
