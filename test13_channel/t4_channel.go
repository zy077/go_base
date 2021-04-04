package main

import "fmt"

func main(){
    // 声明一个channel
    c := make(chan int)

    go func() {
        defer fmt.Println("goroutine end")
        for i:=0; i<5; i++{
            c <- i
        }
        close(c)  // 关闭channel
    }()

    /*for {
        // ok如果为true表示channel没有关闭，如果为false表示channel已经关闭
        if data, ok := <-c; ok{
            fmt.Println(data)
        } else {
            break
        }
    }*/

    // 可以使用range不断迭代从channel中取出数据
    for data := range c {
        fmt.Println(data)
    }

    fmt.Println("main end")
}