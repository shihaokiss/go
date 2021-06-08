package main

import "fmt"

type Book struct {
	name string
}

func Show(arg interface{}) {
	fmt.Println("Show...", arg)
	value_string, is_string := arg.(string)
	if is_string {
		fmt.Println("String do...", value_string)
	}
	value_int, is_int := arg.(int)
	if is_int {
		fmt.Println("Int do...", value_int)
	}
	value_book, is_book := arg.(Book)
	if is_book {
		fmt.Println("Book do...", value_book)
	}
}

func main() {
	book := Book{"aaa"}

	Show("xxx")
	Show(10)
	Show(book)
}
