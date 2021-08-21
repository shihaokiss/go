package main

import "fmt"

type animal interface {
	Move()
	Sleep(int)
}

type Cat struct {
	name string
	age int
}

func (c *Cat) Move() {
    fmt.Println(c.name, c.age, " Move")
}

func (c *Cat) Sleep(hour int) {
	fmt.Println(c.name, c.age, " Sleep")
}

func test() {
	fmt.Println("test::func")
	var a animal
	cat1 := Cat{"cat1", 1}
	cat2 := Cat{"cat2", 2}

	a = cat1
    fmt.Println(a)
	a.Move()
	a.Sleep(20)
	a = &cat2
	fmt.Println(a)
	a.Move()
	a.Sleep(30)
}

func main() {
	test()
}
