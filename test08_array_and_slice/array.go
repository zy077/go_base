package main

import (
	"fmt"
)

func array(){
	// 定义一个长度为10的静态数组，默认值为0
	var array01 [10]int
	for i:=0; i < len(array01); i++ {
		fmt.Println(array01[i])
	}
	fmt.Println("*****************")
	// 定义一个静态数组
	array02 := [10]int{1, 2, 3}
	for idx, val := range array02 {
		fmt.Println(idx, val)
	}
	fmt.Println("+++++++++++++++++++")
}

func printArray(array [10]int){
	array[0] = 110
	// 静态数组作为参数，为值传递，且参数类型一致
	for _, val := range array{  // _:表示匿名变量，只用来占位，不参与程序运行
		fmt.Println(val)
	}
	fmt.Println("++++++++++++++++++++")
}


func main(){
	array()
	// 静态数组作为参数
	array03 := [10]int{1,2,3,4,5}
	printArray(array03)
	fmt.Println("------------------------")
	for _, val := range array03{
		fmt.Println(val)
	}
}
