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

// 数组求和
func ArraySum(array [10]int){
	sum01 := 0
	for i:=0; i<len(array); i++ {
		sum01 += array[i]
	}
	fmt.Println(sum01)
	fmt.Println("++++++++++++++++++++")
	var sum02 int
	for _, item := range array {
		sum02 += item
	}
	fmt.Println(sum02)
	fmt.Println("++++++++++++++++++++")
}

// 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
func FindIndex(array [5]int, num int) []string {
	var ret []string
	for idx01, val01 := range array {
		for idx02, val02 := range array[idx01+1:]{
			if val01 + val02 == num {
				ret = append(ret, fmt.Sprintf("(%d, %d)", idx01, idx01+idx02+1))
			}
		}
	}
	return ret
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
	// 数组求和
	ArraySum(array03)
	array04 := [...]int{1, 3, 5, 7, 8}
	ret := FindIndex(array04, 8)
	fmt.Println(ret)
}
