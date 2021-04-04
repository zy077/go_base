package main

import (
    "fmt"
    "net"
    "sync"
)

// 声明一个Server结构体
type Server struct {
    Ip          string
    Port        int
    OnlineMap   map[string]*User        // 用户在线列表
    mapLock     sync.RWMutex            // 锁
    Message     chan string             // 	//消息广播的channel
}

//监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线User
func (this *Server) ListenServerMessage(){
    for {
        msg := <-this.Message
        // 将消息发送给所有的在线用户
        this.mapLock.Lock()
        for _, cli := range this.OnlineMap {
            cli.UserChan <- msg
        }
        this.mapLock.Unlock()
    }
}


// 广播用户上线消息
func (this *Server) BroadCast(user *User, msg string){
    sendMsg := "[" + user.UserAddress + "]-(" + user.UserName + ")-{" + msg + "}"
    this.Message <-sendMsg
}


func (this *Server) Handler(conn net.Conn){
    // 处理事务
    //fmt.Println("连接建立成功！")

    // 用户上线
    user := NewUser(conn)

    // 将上线用户添加到OnLineMap中
    this.mapLock.Lock()
    this.OnlineMap[user.UserName] = user
    this.mapLock.Unlock()

    // 广播当前用户上线消息
    this.BroadCast(user, "上线啦！")

    // 当前handler阻塞
    select {

    }
}


// 定义服务启动方法
func (this *Server) Start () {
    // 创建一个服务对象
    listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
    if err != nil {
        fmt.Println("Net.Listens err: ", err)
    }
    defer listener.Close()

    //启动监听Message的goroutine
    go this.ListenServerMessage()

    // 监控外界的链接
    // 处理链接
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Listener.Accept err: ", err)
            continue
        }

        // 创建一个go程处理需求
        go this.Handler(conn)
    }
}


// 对外提供一个方法，用于初始化Server对象
func NewServer(ip string, port int) *Server {
    server := Server{
        Ip: ip,
        Port: port,
        OnlineMap: make(map [string]*User),
        Message: make(chan string),
    }
    return &server
}