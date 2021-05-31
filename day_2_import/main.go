package main

import (
	"fmt"
	_ "time"  //不使用
	mylib1 "gobase/day_2_import/lib1"  //重命名
)

func init() {
	fmt.Println("main.init()...")
}

func main() {
	fmt.Println("main.main()...")
	mylib1.Test1()
}
