package main

import "fmt"

func changeValue(a int) {
	a = a * 10
	fmt.Println("changeValue a =", a)
}

func changeValue2(a *int) {
	fmt.Println("changeValue2 a = ", a, *a)
	*a = *a * 10
}

func swap_value(a, b *int) {
	var temp int = 100
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	// aa := 100
	// changeValue(aa)
	// fmt.Println("aa = ", aa)
	// changeValue2(&aa)
	// fmt.Println("aa = ", aa)

	// a, b := 10, 20
	// swap_value(&a, &b)
	// fmt.Println("a = ", a, "\nb = ", b)

	var a int = 100
	var p *int = &a
	var pp **int = &p
	fmt.Println(&a, p, a, *p)
	fmt.Println(&p, pp, p, *pp, **pp)
}
