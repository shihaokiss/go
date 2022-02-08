package main

import "fmt"

type Book struct {
	title string
	auth  string
}

func changeBook1(book Book) {
    book.title = "xxx"
}

func changeBook2(book *Book) {
    book.title = "xxx"
}

type People struct {
	name string
	age int
}

func (this *People) Show() {
	fmt.Println(this.name, this.age)
}

func (this *People) SetName(name string) {
	this.name = name
}

func main() {
	// var book Book
	// book.title = "aaa"
	// book.auth = "bbb"

	// fmt.Println(book)
	// changeBook1(book)
	// fmt.Println(book)
	// changeBook2(&book)
	// fmt.Println(book)

	people := People{"zhangsan", 20}
	people.Show()
	people.SetName("enen")
	people.Show()
}