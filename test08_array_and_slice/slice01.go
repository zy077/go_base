package main

import "fmt"

// 遍历（动态）数组
func test01_slice(array []int)  {
	array[0] = 110
	for _, val := range array{
		fmt.Println(val)
	}
	fmt.Println("++++++++++++++++++")
}

func main()  {
	// 定义一个动态数组
	array01 := []int{1,2,3,4}
	test01_slice(array01)  // 动态数组作为参数传递时传递的是引用

	for _, val := range array01{
		fmt.Println(val)
	}
}