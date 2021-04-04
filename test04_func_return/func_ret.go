package main

import "fmt"

func f01(name string, age int) int {  // 一个返回值
	fmt.Println("name = ", name)
	fmt.Println("age = ", age)
	return 100
}

func f02(a string, b int) (int, int) {  // 多个匿名返回值
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 10, 20
}

func f03(a int, b int) (ret01 int, ret02 int) {  // 多个有名返回值
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	// ret01, ret02相当于f03的形参，默认值为0
	// 作用域为f03{}中的空间
	ret01 = 10
	ret02 = 20
	//return
	return ret01, ret02
}

func main(){
	c := f01("tongtianzi", 20)
	fmt.Println("c = ", c)
	d, e := f02("haha", 10)
	fmt.Println("d = ", d, "e = ", e)
	g, h := f03(100, 200)
	fmt.Println("g = ", g)
	fmt.Println("h = ", h)
}
