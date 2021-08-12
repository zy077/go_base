package main

import "fmt"

// 断言
// 1. 接口类型可以存储任何类型的值
// 2. 我们想要判断其保存的具体是那种类型的值，可以通过断言判断
// 3. 格式：v, ok := x.(T)
//    x:interface类型的变量，断言的类型，v：x转换为T类型后的值，ok：断言是否成功

func JustType(x interface{})  {
    switch v := x.(type) {
    case int:
        fmt.Printf("%d 是int类型 \n", v)
    case string:
        fmt.Printf("%s 是string类型 \n", v)
    default:
        fmt.Printf("Unsupport type \n")
    }
}

func main() {
    x01 := 10
    JustType(x01)
}