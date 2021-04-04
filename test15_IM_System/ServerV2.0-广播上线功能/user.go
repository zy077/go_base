package main

import "net"

//定义User
type User struct {
    UserName            string          // 姓名
    UserAddress         string          // 用户绑定的地址
    UserChan        chan string     // 用户通信的channel
    UserConn        net.Conn        // 用户用于给客户端发送消息的连接
}


// 监听当前User channel的 方法,一旦有消息，就直接发送给对端客户端
func (this *User) ListenMessage(){
    for {
        msg := <-this.UserChan
        this.UserConn.Write([]byte(msg + "\n"))
    }
}


// 创建一个用户的API
func NewUser(conn net.Conn) *User {
    Address := conn.RemoteAddr().String()
    user := &User{
        UserName: Address,
        UserAddress: Address,
        UserChan: make(chan string),
        UserConn: conn,
    }

    // 当这个用户被创建的时就应该启动该用户的goroutine
    go user.ListenMessage()

    return user
}