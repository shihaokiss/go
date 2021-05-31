package main

import "fmt"

func test() {
	const a int = 10
	const (
		aa = 10
		bb = 20
	)
	fmt.Println("a = ", a, "aa = ", aa, "bb = ", bb)
}

// func test() {
// 	var a int
// 	var b int = 100
// 	var c = 100
// 	d := 100

// 	var xx, yy int = 100, 200
// 	var kk, ll = 100, "AceId"
// 	var (
// 		vv int  = 100
// 		jj bool = true
// 	)

// 	fmt.Println("a=", a)
// 	fmt.Println("b=", b)
// 	fmt.Println("c=", c)
// 	fmt.Println("d=", d)

// 	fmt.Println("xx=", xx)
// 	fmt.Println("yy=", yy)
// 	fmt.Println("kk=", kk)
// 	fmt.Println("ll=", ll)
// 	fmt.Println("vv=", vv)
// 	fmt.Println("jj=", jj)
// }

func main() {
	fmt.Println("Hello go.")
	test()
}
