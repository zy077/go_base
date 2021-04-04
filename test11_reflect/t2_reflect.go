package main

import (
	"fmt"
	"reflect"
)

func reflectFunc(args interface{}){
	// 获取args的类型
	fmt.Println(reflect.TypeOf(args))
	// 获取args的值
	fmt.Println(reflect.ValueOf(args))
}

func main()  {
	a := 100
	reflectFunc(a)
	var b string = "haha"
	reflectFunc(b)
}