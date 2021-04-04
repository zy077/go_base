package main

import (
	"fmt"
)

/*
四种变量声明方式
*/

// 声明全局变量（方法4不能用来声明全局变量）
var g01 int
var g02 string = "tongtianzi"

func main() {
	// 第一种：只声明，不赋值，默认值为0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("a type is: %T\n", a)

	// 第二种：声明并赋值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("b type is: %T\n", b)

	// 第三种：初始化时省去数据类型，通过自动匹配当前变量的数据类型
	var c = "haha"
	fmt.Println("c = ", c)
	fmt.Printf("c type is: %T\n", c)

	// 第四种（最常用）：省去var关键字和数据类型
	d := 200
	fmt.Println("d = ", d)
	fmt.Printf("d type is: %T\n", d)

	// 访问全局变量
	fmt.Println("g01 = ", g01, "g02 = ", g02)

	// 声明多个变量
	var e, f int = 101, 102
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
	var g, h = 201, 202
	fmt.Println("g = ", g, "h = ", h)

	var (
		j int = 1001
		k string = "TongTianZi"
	)
	fmt.Println("j = ", j)
	fmt.Println("k = ", k)
}
