package main

import "fmt"

// const来定义枚举
const (
	// 在const()中通过关键字iota来实现自增枚举，每行iota加1，iota默认为0
	a = iota  // iota = 0
	b         // iota = 1
)

const (
	c, d = iota + 1, iota + 2  // iota = 0, c = 1, d = 2
	e, f 						// iota = 1, e = 2, f = 3
	g, h						// iota = 2, g = 3, h = 4

	j, k = 10 * iota, 20 * iota  // iota = 3, j  = 30, k = 60
	l, m						// iota = 4, l = 40, m = 80
)

func main(){
	// const定义常量
	const height int = 100
	fmt.Println("height = ", height)

	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(e, f)
	fmt.Println(g, h)
	fmt.Println(j, k)
	fmt.Println(l, m)
}
