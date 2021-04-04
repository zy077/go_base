package main

import (
	"fmt"
)

// interface{}：通用万能类型
func test_interface(args interface{}){
	fmt.Println(args)

	// 通过断言判断args到底是什么类型
	val, isString := args.(string)  // 判断args是否是string类型
	if isString {
		fmt.Printf("%s type is %T\n", val, val)
	}else {
		fmt.Printf("%s type is %T\n", val, val)
	}
	fmt.Println("==================")
}

type Book01 struct {
	BookName string
	Author string
}

func main()  {
	// 万能类型测试
	test_interface(10)
	test_interface("haha")

	b := Book01{BookName: "傻吊英雄传", Author: "金庸"}
	test_interface(b)
}