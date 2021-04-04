package main

import (
    "fmt"
    "time"
)


func WriteInto(c01, c02 chan int) {
    fmt.Println("开始写入数据...")
    for i:=0; i<10; i++{
        c01 <- i
    }
    time.Sleep(1 * time.Second)
    c02 <- 10
}

func main(){
    // 声明两个channel
    c01 := make(chan int)
    c02 := make(chan int)

    // 开启一个协成，往c01中写数据
    go WriteInto(c01, c02)

    for {
        select {
            case <-c01:  // 监控c01中是否有数据，如果有数据，就从c01中取出来，添加到c02中
                fmt.Println(<-c01)
            case <-c02:
                return
        }
    }

    fmt.Println("main end ...")
}
