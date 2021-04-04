package main

import (
	"fmt"
	"time"
)

func GoroutineTest(){
    for {  // 不写控制条件就是死循环
        fmt.Println("Goroutine running...")
        time.Sleep(1 * time.Second)
    }
}

func main(){
    // 协成（goroutine）：又名微线程，相比于线程，消耗资源更少
    go GoroutineTest()  // 使用go关键字开启一个协成

    for {
        fmt.Println("main running...")
        time.Sleep(1 * time.Second)
    }
}
