package main

import "fmt"

// 匿名函数
func main() {
    // 将匿名函数保存到变量
    add := func(x, y int) int {
        return x + y
    }
    ret := add(10, 20)  // 通过变量调用匿名函数
    fmt.Println(ret)

    //自执行函数：匿名函数定义完加()直接执行
    func(a, b int) {
        fmt.Printf("a = %d, b = %d", a, b)
    }(100, 200)
}

// 匿名函数多用于实现匿名函数和闭包