package main

import (
	"fmt"
	"io"
	"os"
)

// 变量对应的 pair<type, value>
func test() {
	// pair: <type:*os.File, value:"/dev/tty"文件描述符>
	tty, err := os.OpenFile("test.go", os.O_RDWR, 0)
	println("tty err", tty, err)

	if err != nil {
		fmt.Printf("io.Reader::%T\n", err)
	}

	// pair: <type:, value:>
	var r io.Reader
	fmt.Printf("io.Reader::%T\n", r)
	// pair: <type:*os.File, value:"/dev/tty"文件描述符>
	r = tty
	fmt.Printf("io.Reader::%T\n", r)

	// pair: <type:, value:>
	var ww io.Writer
	fmt.Printf("io.Writer::%T\n", ww)
	// pair: <type:*os.File, value:"/dev/tty"文件描述符>
	ww = r.(io.Writer) // 强制转换
	fmt.Printf("io.Writer::%T\n", ww)

	ww.Write([]byte("写向终端内容aaa！！\n"))

}

type ReadBook interface {
	ReadBook()
}
type WriteBook interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("你好 我要阅读一本书！！！")
}
func (this *Book) WriteBook() {
	fmt.Println("你好 我要写作一本书！！！")
}

func test2() {
	// pair<type:&Book, value:Book地址>
	b := &Book{}
	fmt.Println("Book b: ", b)
	println(b)
	b.ReadBook()
	b.WriteBook()

	var r ReadBook
	r = b
	r.ReadBook()

	var w WriteBook
	w = r.(WriteBook)
	w.WriteBook()

}

func main() {
	test2()
}
