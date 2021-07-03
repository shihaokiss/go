package main

import "fmt"

func TestPair() {
	var nnn string = "xxx"
	fmt.Println(nnn)

	name := "haha"
	var all_type interface{} = name
	fmt.Println(name, all_type)

	str, is_str := all_type.(string)
	fmt.Println(str, is_str)
}

func main() {
	TestPair()
}
