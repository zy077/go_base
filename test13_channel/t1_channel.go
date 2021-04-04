package main

import "fmt"

func main(){
    // 声明一个channel
    c01 := make(chan int)

    // goroutine1
    go func(){
        defer fmt.Println("goroutine1 end ...")
        fmt.Println("goroutine1 start ...")
        c01 <- 666 // 将666发送到c01中
    }()

    // goroutine2
    go func(){
        defer fmt.Println("goroutine2 end ...")
        fmt.Println("goroutine2 start ...")
        c02 := <- c01  // 从c01中读取666
        fmt.Println(c02)
    }()

    fmt.Println("main end ...")
}
