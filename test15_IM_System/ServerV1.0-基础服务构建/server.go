package main

import (
    "fmt"
    "net"
)

// 声明一个Server结构体
type Server struct {
    Ip     string
    Port   int
}


func (this *Server) Handler(conn net.Conn){
    // 处理事务
    fmt.Println("连接建立成功！")
}


// 定义服务启动方法
func (this *Server) Start () {
    // 创建一个服务对象
    listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
    if err != nil {
        fmt.Println("Net.Listens err: ", err)
    }
    defer listener.Close()

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
    }
    return &server
}