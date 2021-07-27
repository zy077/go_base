package main

import (
	"fmt"
	"sort"
)

// 遍历（动态）数组
func test01_slice(array []int)  {
	array[0] = 110
	for _, val := range array{
		fmt.Println(val)
	}
	fmt.Println("++++++++++++++++++")
}

// 切片排序
func SliceSort(slice0 []int) []int {
	sort.Ints(slice0)
	return slice0
}

func main()  {
	// 定义一个动态数组
	array01 := []int{1,2,3,4}
	test01_slice(array01)  // 动态数组作为参数传递时传递的是引用

	for _, val := range array01{
		fmt.Println(val)
	}

	slice01 := make([]int, 10, 10)
	slice02 := slice01[4:]
	fmt.Printf("len: %d, cap: %d", len(slice02), len(slice02))
	fmt.Println("++++++++++++++++++")

	array02 := [6]int{1,2,3,4,5,6}
	slice03 := array02[2:]
	fmt.Printf("len: %d, cap: %d", len(slice03), len(slice03))
	fmt.Println("++++++++++++++++++")

	a := make([]string, 5, 10)
	for i:=0; i<10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
	fmt.Println("++++++++++++++++++")

	slice04 := []int{3, 7, 8, 9, 1}
	slice05 := SliceSort(slice04)
	fmt.Println(slice05)
	fmt.Println("++++++++++++++++++")
}