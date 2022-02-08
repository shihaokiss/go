package main

import (
	"encoding/json"
	"fmt"
)

// type animal interface {
// 	Move()
// 	Sleep(int)
// }

// type Cat struct {
// 	name string
// 	age int
// }

// func (c *Cat) Move() {
//     fmt.Println(c.name, c.age, " Move")
// }

// func (c *Cat) Sleep(hour int) {
// 	fmt.Println(c.name, c.age, " Sleep")
// }

// func test() {
// 	fmt.Println("test::func")
// 	var a animal
// 	cat1 := Cat{"cat1", 1}
// 	cat2 := Cat{"cat2", 2}

// 	a = cat1
//     fmt.Println(a)
// 	a.Move()
// 	a.Sleep(20)
// 	a = &cat2
// 	fmt.Println(a)
// 	a.Move()
// 	a.Sleep(30)
// }

type student struct {
    Name string `json:"name"`
	Age int `json:"age"`
}

func test() {
    var stu student
	student_str := `{"name": "zhangsan", "age": 200}`
	json.Unmarshal([]byte(student_str), &stu)

	fmt.Println(stu.Name, stu.Age)
}

func main() {
	test()
}
