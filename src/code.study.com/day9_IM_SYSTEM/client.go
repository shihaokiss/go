package main

import (
	"flag"
	"fmt"
	"io"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	Conn       net.Conn
	flag       int
}

func (this *Client) Menu() bool {
	fmt.Println("1. 公聊模式")
	fmt.Println("2. 私聊模式")
	fmt.Println("3. 修改用户名")
	fmt.Println("4. 查询用户")
	fmt.Println("0. 退出")

	var flag int
	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 4 {
		this.flag = flag
		return true
	} else {
		fmt.Println("input flag error.")
		return false
	}
}

//公聊模式
func (this *Client) PublicChat() {
	fmt.Println(">>>>>>please input message (enter exit to exit): ")
	var sendMsg string
	fmt.Scanln(&sendMsg)
    for sendMsg != "exit" {
        if len(sendMsg) > 0 {
			sendMsg = sendMsg + "\n"
			_, err := this.Conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("PublicChat to server err: ", err)
				return
			}
		}
		fmt.Println(">>>>>>please input message (enter exit to exit): ")
        sendMsg = ""
	    fmt.Scanln(&sendMsg)
	}
	fmt.Println(">>>>>>PublicChat exit.")
}

//私聊模式
func (this *Client) PrivateChat() {
	this.UserInfo()
    fmt.Println(">>>>>>please select user (enter exit to exit): ")
	var userName string
	fmt.Scanln(&userName)
	for userName != "exit" {
        fmt.Println(">>>>>>please input msg (enter exit to exit): ")
		var sendMsg string
		fmt.Scanln(&sendMsg)
		for sendMsg != "exit" {
			if len(sendMsg) > 0 {
				sendMsg = "to|" + userName + "|" + sendMsg + "\n"
				_, err := this.Conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("PublicChat to server err: ", err)
				    return
				}
			}
			fmt.Println(">>>>>>pleast input msg (enter exit to exit): ")
			sendMsg = ""
			fmt.Scanln(&sendMsg)
		}
        fmt.Println(">>>>>>please select user (enter exit to exit): ")
		userName = ""
		fmt.Scanln(&userName)
	}
	fmt.Println(">>>>>>PrivateChat exit.")
}

func (this *Client) ReName() bool {
	fmt.Println("please input new username:")
	var newName string
	fmt.Scanln(&newName)

	sendMsg := "rename|" + newName + "\n"
	_, err := this.Conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("send rename msg to server err: ", err)
		return false
	}
	return true
}

func (this *Client) UserInfo() bool {
	sendMsg := "who\n"
	_, err := this.Conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("send rename msg to server err: ", err)
		return false
	}
	return true
}

func (this *Client) Run() {
	for this.flag != 0 {
		if this.Menu() != true {
			continue
		}

		//根据不同的业务输出不同的逻辑
		switch this.flag {
		case 1:
			this.PublicChat()
			break
		case 2:
			this.PrivateChat()
			break
		case 3:
			this.ReName()
			break
		case 4:
			this.UserInfo()
			break
		}
	}
	fmt.Println("exit bye ~")
}

func (this *Client) DoMessage(msg string) {
	fmt.Println(msg)
}

func (this *Client) DealResponse() {
	// 只要 this.Conn 有数据，就会显示到标准输出上面
	// io.Copy(os.Stdout, this.Conn)
	// 等同于上面的逻辑
	for {
		buf := make([]byte, 4096)
		n, err := this.Conn.Read(buf)
        if n == 0 {
            fmt.Println("client-server close.")
			return
		}
		if err != nil && err != io.EOF {
			fmt.Println("client read err: ", err)
			return
		}
		msg := string(buf[:n-1])
		//客户端处理消息
		this.DoMessage(msg)
	}
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", client.ServerIp, client.ServerPort))
	if err != nil {
		fmt.Printf("connect to server err: ", err)
		return nil
	}
	client.Conn = conn
	//返回client对象
	return client
}

//////////////////////////////////////////////////////////////
var ServerIp string
var ServerPort int

func init() {
	flag.StringVar(&ServerIp, "ip", "127.0.0.1", "服务端IP地址(默认为127.0.0.1)")
	flag.IntVar(&ServerPort, "port", 8888, "服务端PORT(默认为8888)")
}

func main() {
	//参数解析
	flag.Parse()

	client := NewClient(ServerIp, ServerPort)
	if client == nil {
		fmt.Println(">>>>>>>> Client to Server fail.")
	} else {
		fmt.Println(">>>>>>>> Client to Server success.")
	}

	//处理服务端响应的消息
	go client.DealResponse()

	//使得主程序不退出
	client.Run()
}
