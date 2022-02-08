package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip         string
	Port       int
	mapLock    sync.RWMutex
	OnlineUser map[string]*User
	Message    chan string
}

//创建一个服务对象
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:         ip,
		Port:       port,
		OnlineUser: make(map[string]*User),
		Message:    make(chan string),
	}
	return server
}

//启动服务器的接口, Server 类的一个方法
func (this *Server) Start() {
	//listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	// 函数推出时候关闭 close
	defer listener.Close()
	go this.ListenMessageChannel()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.accept err:", err)
			continue
		}

		//dohandle
		go this.DoHandle(conn)
	}
}

// 向message管道中写信息,并广播
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + "-" + user.Name + "-" + msg
	this.Message <- sendMsg
}

// 监听message 管道的goroutine, 把消息分发到各个用户
func (this *Server) ListenMessageChannel() {
	for {
		msg := <-this.Message
		this.mapLock.Lock()
		for _, user := range this.OnlineUser {
			user.Channel <- msg
		}
		this.mapLock.Unlock()
	}
}

// Server dohandle 的方法
func (this *Server) DoHandle(conn net.Conn) {
	fmt.Println("do handle start...")
	user := NewUser(conn, this)
	user.Online()

	//监听用户是否活跃的channel
	isLive := make(chan bool)

	//接受客户端的消息并进行广播
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				//用户下线
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println(user.Addr, " read error:", err)
				return
			}
			//提取用户的消息,去除 '\n'
			msg := string(buf[:n-1])

			//用户处理消息
			user.DoMessage(msg)
			isLive <- true
		}
	}()

	//handle 阻塞
	for {
		select {
		case <-isLive:
			//用户活跃，检测time.After的时候会归0触发器
		case <-time.After(time.Second * 60):
			//超时了，需要踢出用户
			user.SendMsg("timeout!")
			user.Clean()
			//清理完成不用走用户下线的逻辑，直接退出DoHandle
			return
		}
	}
}
