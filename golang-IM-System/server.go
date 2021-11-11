package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip string
	Port int

	//在线用户的列表
	OnlineMap map[string]*User
	maplock sync.RWMutex

	//消息广播的channel
	Message chan string
}

//创建一个对外的接口
func NewServer(ip string,port int) *Server {
	server := &Server{
		Ip : ip,
		Port : port,
		OnlineMap: make(map[string]*User),
		Message: make(chan string),
	}
	return server
}

//监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线user
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		//将message发送给所有在线User
		this.maplock.Lock()
		for _,cil := range this.OnlineMap {
			cil.C <- msg
		}
		this.maplock.Unlock()
	}
}

//广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMessage := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- sendMessage
}

func (this *Server) Handler(conn net.Conn) {
	//当前处理的业务
	//fmt.Println("链接建立成功")

	user := NewUser(conn, this)

	user.Online()

	//监听用户是否活跃的管道
	isLive := make(chan bool)

	//接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read Error",err)
				return
			}

			//提取用户的消息(去除‘\n’)
			msg := string(buf[:n-1])

			//将得到的消息进行广播
			user.DoMessage(msg)

			isLive <- true
		}
	}()
	//判断用户是否活跃，不活跃强踢
	for {
		select {
		case <-isLive:
			//什么也不做，为了激活下面的计时器判断，以重置计时器
		case <-time.After(time.Second * 120):
			//已经超时
			user.SendMessage("You are fired")

			//销毁该用户的资源
			close(user.C)

			//关闭连接
			conn.Close()

			//退出当前的Handler
			return
		}
	}
}

//启动服务器的接口
func (this *Server) Start() {
	//socket listen
	listener,err := net.Listen("tcp",fmt.Sprintf("%s:%d",this.Ip,this.Port))
	if err != nil {
		fmt.Println("net.listen error",err)
		return
	}
	//close listen socket
	defer listener.Close()

	//启动监听Message的goroutine
	go this.ListenMessager()

	for {
		//accept
		conn ,err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error",err)
			continue
		}

		//do handler
		go this.Handler(conn)
	}
}