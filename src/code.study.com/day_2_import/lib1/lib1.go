package lib1

import (
	"fmt"
	"gobase/day_2_import/lib2"
)

func Test1() {
	fmt.Println("lib1.Test()...")
	lib2.Test2()
}

func init() {
	fmt.Println("lib1.init()...")
}
