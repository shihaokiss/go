package main

import (
	"fmt"
	"net"
	"strings"
)


type User struct {
	Name    string
	Addr    string
	Channel chan string
	Conn    net.Conn
	Server  *Server
}

//创建一个用户信息
func NewUser(conn net.Conn, server *Server) *User {
	addr := conn.RemoteAddr().String()
	user := &User {
        Name: addr,
		Addr: addr,
		Channel: make(chan string),
		Conn: conn,
		Server: server,
	}

	//启动监听管道中的goroutine(go程)
	fmt.Println("NewUser ListenMessage start.")
	go user.ListenMessage()
	return user
}

//客户端监听管道
func (this *User) ListenMessage() {
	for {
		//监听到管道中有数据，就把信息写到对端
		message := <-this.Channel
		this.SendMsg(message)
	}
}

//用户上线
func (this *User) Online() {
	//用户上线
	this.Server.mapLock.Lock()
	this.Server.OnlineUser[this.Name] = this
	this.Server.mapLock.Unlock()
    this.Server.BroadCast(this, "online!")
}

//用户下线
func (this *User) Offline() {
	//用户下线
	this.Server.mapLock.Lock()
	delete(this.Server.OnlineUser, this.Name)
	this.Server.mapLock.Unlock()
    this.Server.BroadCast(this, "offline!")
}

//清理用户资源
func (this *User) Clean() {
	close(this.Channel)
	this.Conn.Close()
}

//发送消息
func (this *User) SendMsg(msg string) {
	this.Conn.Write([]byte(msg+"\n"))
}

//用户消息处理
func (this *User) DoMessage(msg string) {
	//服务端进行广播
	if msg == "who" {
        this.Server.mapLock.Lock()
		for _, user := range this.Server.OnlineUser {
			msg := "[" + user.Addr + "]" + user.Name + ": online!"
			this.SendMsg(msg)
		}
		this.Server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := strings.Split(msg, "|")[1]
        // 判断有没有用户重名
		_, ok := this.Server.OnlineUser[newName]
		if ok {
			this.SendMsg(newName + " has already exist!")
		} else {
			this.Server.mapLock.Lock()
			delete(this.Server.OnlineUser, this.Name)
			this.Server.OnlineUser[newName] = this
			this.Server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("rename success! user=" + newName)
		}

	} else if len(msg) > 4 && msg[:3] == "to|" {
        //消息格式 to|name|msg
		userName := strings.Split(msg, "|")[1]
		remoteUser, ok := this.Server.OnlineUser[userName]
		if !ok {
            this.SendMsg("this user not exist!")
			return
		}

		msg := strings.Split(msg, "|")[2]
		if msg == "" {
			this.SendMsg("can not send empty msg!")
			return
		}
		
		remoteUser.SendMsg("[" + this.Name + "] said to you: " + msg)
	} else {
		this.Server.BroadCast(this, msg)
	}
}