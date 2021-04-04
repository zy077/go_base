package main

import "fmt"

// 声明一种新的数据类型myInt，是int的别名
type myInt int


// 定义一个结构体
type Book struct {
	title string
	author string
}


func changeBook0(book Book){
	book.title = "活着"
	book.author = "余华"
	fmt.Printf("书名：%s, 作者：%s\n", book.title, book.author)
	fmt.Println("==========================")
}


func changeBook1(book *Book){
	book.title = "白鹿原"
	book.author = "陈忠实"
	fmt.Printf("书名：%s, 作者：%s\n", book.title, book.author)
	fmt.Println("==========================")
}

func main()  {
	// myInt
	var num01 myInt = 100
	fmt.Println(num01)
	fmt.Printf("%T\n", num01)
	fmt.Println("==========================")

	// 结构体的使用
	var book0 Book
	book0.title = "平凡的世界"
	book0.author = "路遥"
	fmt.Printf("书名：%s, 作者：%s\n", book0.title, book0.author)
	fmt.Println("==========================")

	// 结构体作为参数传递的一个副本（值传递）
	changeBook0(book0)
	fmt.Printf("书名：%s, 作者：%s\n", book0.title, book0.author)
	fmt.Println("==========================")

	// 如果要把结构体的地址值作为参数传递，需要指明形参为指针类型，实参传递地址值
	changeBook1(&book0)
	fmt.Printf("书名：%s, 作者：%s\n", book0.title, book0.author)
	fmt.Println("==========================")
}
