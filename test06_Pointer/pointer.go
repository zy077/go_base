package main

import "fmt"

func swap(pa *int, pb *int){  // 表示pa和pb都是指针类型
	temp := *pa  // 用temp保存pa的地址值
	*pa = *pb  // pa的地址值等于pb的地址值
	*pb = temp  // pb的地址值等于pa的地址值
}

func main() {
	var a = 10
	var b = 20
	// 交换a , b的值
	swap(&a, &b)  // 传递的是a和b的地址值
	fmt.Println("a = ", a, "b = ",  b)
}


