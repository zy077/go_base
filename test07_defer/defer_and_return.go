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

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x  // 5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5  // 6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x  // 5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5  // 5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))  // (A, 1, 2, 3); (AA, 1, 3, 4)
	x = 10
	defer calc("BB", x, calc("B", x, y))  // (B, 10, 2, 12); (BB, 10, 12, 22)
	y = 20

	/*
	A 1 2 3
	B 10 2 12
	BB 10 12 22
	AA 1 3 4
	*/

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}


