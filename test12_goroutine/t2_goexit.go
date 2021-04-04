package main

import (
    "fmt"
    "runtime"
)

func main(){
    // go开启一个协程序
    go func(){  // 定义一个匿名函数
        defer fmt.Println("defer A...")
        func(){
            defer fmt.Println("defer B...")
            // return "B"不会被打印，"A"还是会被打印  相当于continue
            runtime.Goexit()  // 退出当前goroutine，"B"和"A"都不会被打印  相当于break
            fmt.Println("B")
        }()
        fmt.Println("A ")
    }()  // 调用匿名函数

    fmt.Println("main running ...")
}
