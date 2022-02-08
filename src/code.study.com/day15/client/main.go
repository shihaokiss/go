package main

import (
	"fmt"
	"net"
	"time"
)

func client(ip string, port int) {
	// 1. 连接服务端
	// 2. 发送10条信息
	// 3. 关闭socket退出

    conn, error := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if error != nil {
		fmt.Println("client Connect error: e", error)
	}
	defer conn.Close()

	for i:=0; i<20; i++ {
		msg := "nihao wojiao shihao."
		conn.Write([]byte(msg))
		time.Sleep(time.Second)
	}
}

func main() {
	client("127.0.0.1", 8888)
}