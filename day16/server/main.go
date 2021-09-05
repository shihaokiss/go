package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func dohandle(w http.ResponseWriter, r *http.Request) {
    b, err := ioutil.ReadFile("./file.html")
	if err != nil {
		fmt.Println("读文件错误：", err)
	}
	w.Write(b)
}

func server(ip string, port int) {
    // 1. 启动web服务
	// 2. 读 .html 文件并返回
	http.HandleFunc("/readfile/", dohandle)
	http.ListenAndServe(fmt.Sprintf("%s:%d", ip, port), nil)
}

func main() {
    server("127.0.0.1", 9999)
}