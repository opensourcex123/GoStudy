package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

//创建一个客户端
type Client struct {
	ServerIp string
	ServerPort int
	Name string
	conn net.Conn
	flag int	//用户选择菜单的标识
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端对象
	client := &Client{
		ServerIp: serverIp,
		ServerPort: serverPort,
		flag: 999,
	}

	//链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error", err)
		return nil
	}

	client.conn = conn

	//返回客户端
	return client
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器Ip地址（默认是127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口", )
}

//处理server的返回消息，直接显示标准输出
func (client *Client) DealResponse() {
	//一旦client.conn有数据，就直接copy到stdout标准输出中，永久监听阻塞
	io.Copy(os.Stdout, client.conn)
}


func (client *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >=0 && flag <=3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>please input right number")
		return false
	}
}

//查询在线用户
func (client *Client) SelectUsers() {
	sendMsg := "who\n"

	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write error")
		return
	}
}

//私聊模式
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	client.SelectUsers()
	fmt.Println("please input a chat object,input exit to exit")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println("please input chat content,input exit to exit")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn write error", err)
					break
				}
			}
			chatMsg = ""
			fmt.Println("please input chat content,input exit to exit")
			fmt.Scanln(&chatMsg)
		}

		client.SelectUsers()
		fmt.Println("please input a chat object,input exit to exit")
		fmt.Scanln(&remoteName)
	}
}

func (client *Client) GroupChat() {
	//提示用户输入消息
	var chatMsg string

	fmt.Println("please input chat content,input \"exit\" to exit")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		//发送给服务器
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn write error", err)
				break
			}
		}
		chatMsg = ""
		fmt.Println("please input chat content,input \"exit\" to exit")
		fmt.Scanln(&chatMsg)
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println(">>>please input updated name")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("Conn write error", err)
		return false
	}

	return true
}

func (client *Client) Run() {
	for client.flag != 0{
		for client.menu() != true {
		}
		//根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			//公聊模式
			client.GroupChat()
			break
		case 2:
			//私聊模式
			client.PrivateChat()
			break
		case 3:
			//更新用户名
			client.UpdateName()
			break
		}
	}
}

func main() {
	//命令行解析
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>> connection error...")
		return
	}
	fmt.Println("connection succeed")

	//启动一个goroutine，来监听服务器返回的数据
	go client.DealResponse()

	//启动客户端的业务
	client.Run()
}