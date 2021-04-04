package main

import "fmt"

func test01_defer() {
	fmt.Println("test01_defer ...")
}

func test02_defer() {
	fmt.Println("test02_defer ...")
}

func main() {
	defer test01_defer()  // test01_defer进栈
	defer test02_defer()  // test01_defer进栈
	fmt.Println("main ...")
	// 在函数结束前出栈
}
