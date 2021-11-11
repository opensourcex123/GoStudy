package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C chan string
	conn net.Conn

	server *Server //将用户与服务器管关联起来
}

//创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{userAddr,userAddr,make(chan string),conn,server}

	//启动监听当前user，channel消息的goroutine
	go user.ListenMessage()
	return user
}

//监听当前的User Channel的方法，一旦有消息，就直接发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}

//用户的上线业务
func (this *User) Online() {
	//用户上线，将用户加入到onlineMap中
	this.server.maplock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.maplock.Unlock()

	//广播当前用户上线消息
	this.server.BroadCast(this,"online")
}

//用户的下线业务
func (this *User) Offline() {
	//用户下线，将用户从onlineMap中删除
	this.server.maplock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.maplock.Unlock()

	//广播当前用户下线消息
	this.server.BroadCast(this,"offline")
}

//给当前用户发送消息的方法
func (this *User) SendMessage(msg string) {
	this.conn.Write([]byte(msg))
}

//用户的处理消息业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//查询当前在线用户有哪些
		this.server.maplock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMessage := "[" + user.Addr + "]" +user.Name + ":" +"is online...\n"
			this.SendMessage(onlineMessage)
		}
		this.server.maplock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|"{
		//用户修改用户名，格式为rename|名字
		newName := strings.Split(msg, "|")[1]

		//判断name是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMessage("user name is existing")
		} else {
			this.server.maplock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.maplock.Unlock()

			this.Name = newName
			this.SendMessage("Your new name " + this.Name + " is updated successfully" + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|"{
		//私聊功能实现,消息格式为to|用户名|消息
		//1.获取对方的用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMessage("Your format is wrong,please use \"to|john|hello\" format")
			return
		}
		//2.根据用户名，得到user对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMessage("This user doesn't exist")
			return
		}
		//3.获取消息内容，通过对方的user对象把消息发送过去
		remoteMsg := strings.Split(msg, "|")[2]
		if remoteMsg == "" {
			this.SendMessage("The message is blank,please repeat it")
			return
		}
		remoteUser.SendMessage(this.Name + " talk to you that " + remoteMsg)
	} else {
		this.server.BroadCast(this, msg)
	}
}