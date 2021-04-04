package main

import fmt "fmt"

func main() {
	// 声明一个切片，长度为3，容量为5
	var slice01 = make([]int, 3, 5)
	fmt.Printf("length = %d, cap = %d, slice = %v\n", len(slice01), cap(slice01), slice01)
	// 追加一个元素
	slice01 = append(slice01, 10)
	// 追加第二个元素
	slice01 = append(slice01, 20)
	fmt.Printf("length = %d, cap = %d, slice = %v\n", len(slice01), cap(slice01), slice01)
	// 追加第三个元素
	slice01 = append(slice01, 30)
	fmt.Printf("length = %d, cap = %d, slice = %v\n", len(slice01), cap(slice01), slice01)
	fmt.Println("++++++++++++++++++++++++++++")

	var slice02 = make([]int, 2)  // 如果不指定容量，容量默认和长度相同
	fmt.Printf("length = %d, cap = %d, slice = %v\n", len(slice02), cap(slice02), slice02)
	slice02 = append(slice02, 100)
	fmt.Printf("length = %d, cap = %d, slice = %v\n", len(slice02), cap(slice02), slice02)
	slice02 = append(slice02, 101)
	fmt.Printf("length = %d, cap = %d, slice = %v\n", len(slice02), cap(slice02), slice02)
	slice02 = append(slice02, 102)
	fmt.Printf("length = %d, cap = %d, slice = %v\n", len(slice02), cap(slice02), slice02)
}
