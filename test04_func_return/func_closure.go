package main

import "fmt"

// 闭包
// 理解闭包最方便的方法就是将闭包函数看成一个类

// 示例1
func Add01() func(int, int) int {
    z := 10
    return func(x int, y int) int {
        return x + y + z
    }
}

// 示例2
func Add02(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}

// 装饰器和中间件指的就是函数作为参数进行传递的情况
func wrapping(f func() string) {
    fmt.Println("Before working...")
    fmt.Println("Working: ", f())
    fmt.Println("After working...")
}

func sayHello() string {
    return "Hello !"
}

func sayByeBye() string {
    return "Bye Bye ！"
}

func main()  {
    f := Add01()
    fmt.Println(f(20, 30))
    f02 := Add02(10)
    fmt.Println(f02(20))

    // 装饰器和中间件
    wrapping(sayHello)
    wrapping(sayByeBye)
}