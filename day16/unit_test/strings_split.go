package main

import (
	// "fmt"
	"flag"
	"fmt"
	"strings"
)

// "abcaacdd" "b"
func Split(str string, sep string) []string {
	res := make([]string, 0, strings.Count(str, sep)+1)
	index := strings.Index(str, sep)
	for index > -1 {
		res = append(res, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	res = append(res, str)
	return res
}

func Flag() {
	name := flag.String("name", "王大锤", "请输入你是什么大锤")
	flag.Parse()
    fmt.Println(*name)
}

func main() {
	Flag()
}
