package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func do_handle(conn net.Conn) {
    defer conn.Close()
	reader := bufio.NewReader(conn)
    
	var buf [1024]byte
	for {
        n, err := reader.Read(buf[:])
		if err == io.EOF {
			// 读到文件末尾了
            break
		}
		if err != nil {
			fmt.Println("do_handle read err: ", err)
		}

		msg := string(buf[:n])
		fmt.Println("服务端读到：", msg)
	}
}

func server(ip string, port int) {
	// 1. 监听端口
	// 2. 接受消息
	// 3. 处理消息
	// 4. 返回消息
	// 5. 关闭连接

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("server connect err: ", err)
		return
	}
	defer listener.Close()
	fmt.Println("server listen:", port)
    
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err: ", err)
			return
		}

		go do_handle(conn)
	}
}

func main() {
	server("127.0.0.1", 8888)
}