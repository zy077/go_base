package main

import "fmt"

func main() {
	// 切片的四种声明方式
	// 方式一：声明slice1是一个切片，并且初始化，默认值是1，2，3。 长度len是3
	slice01 := []int{1, 2, 3}
	fmt.Println(slice01)
	fmt.Println("+++++++++++++++++++++")

	// 方式二：
	var slice02 []int  // 声明slice02是一个切片，并没有给slice02分配空间
	slice02 = make([]int, 3)  // 给slice02开辟3个空间，默认值为0
	fmt.Println(slice02)
	fmt.Println("+++++++++++++++++++++")

	// 方式三：声明slice03是一个切片，并且开辟3个空间，默认值为0
	var slice03 []int = make([]int, 3)
	fmt.Println(slice03)
	fmt.Println("+++++++++++++++++++++")

	// 方式四(常用)：声明slice1是一个切片，同时给slice分配空间，3个空间，初始化值是0, 通过:=推导出slice是一个切片
	slice04 := make([]int, 3)
	fmt.Println(slice04)
	fmt.Println("+++++++++++++++++++++")

	if slice04 == nil {
		fmt.Println("slice04 为空")
	}else {
		fmt.Println("slice04 不为空")
	}
}
