package main

import "fmt"

func returnFunc() int {
	fmt.Println("return 执行了...")
	return 0
}

func deferFunc() {
	fmt.Println("defer 执行了...")
}

func return_defer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	return_defer()
}
