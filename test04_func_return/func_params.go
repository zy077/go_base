package main

import "fmt"

func Sum(l... int) int {
    fmt.Println(l)  // 切片
    fmt.Printf("%T \n", l)
    ret := 0
    for _, val := range l {
        ret += val
    }
    return ret
}

//固定参数搭配可变参数使用时，可变参数要放在固定参数的后面
func Sum01(x int, l... int) int {
    fmt.Printf("%T \n", x)
    fmt.Println(x)
    fmt.Printf("%T \n", l)
    fmt.Println(l)
    for _, val := range l {
        x += val
    }
    return x
}


func main()  {
    fmt.Println(Sum(10, 20, 30))
    fmt.Println(Sum01(10, 1, 2))
}