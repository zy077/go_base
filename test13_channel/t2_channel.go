package main

import "fmt"

func main(){
    c01 := make(chan int, 3) // 初始化带有缓冲的channel
    fmt.Println("c01 = ", c01, ", len() = ", len(c01), ", cap() = ", cap(c01))

    // write
    go func() {
        defer fmt.Println("write end")
        for i:=0; i < cap(c01); i++ {
            c01 <- i
            fmt.Println("write in ", i)
        }
    }()

    // read
    go func() {
        defer fmt.Println("read end")
        for i:=0; i<cap(c01); i++ {
            c02 := <-c01
            fmt.Println("read in ", c02)
        }
    }()

    fmt.Println("main end")
}

