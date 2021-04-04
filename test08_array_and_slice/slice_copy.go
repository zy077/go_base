package main

import "fmt"

func main()  {
	// 截取（即python中的切片）：切片是一种引用
	s0 := []int{1, 2, 3}  // [1, 2, 3]
	fmt.Println(s0)
	s1 := s0[0:2]  // [1, 2]
	fmt.Println(s1)
	// 修改
	s0[0] = 100  // [100, 2, 3]
	fmt.Println(s0)  // [100, 2, 3]
	fmt.Println(s1)  // [100, 2]
	fmt.Println("===========================")

	// copy
	s3 := make([]int, 3)  // [0, 0, 0]
	copy(s3, s0) // 将s0中的值 依次拷贝到s3中
	fmt.Println(s3)  // [100, 2, 3]
}
