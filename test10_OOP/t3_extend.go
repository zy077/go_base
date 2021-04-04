package main

import (
	"fmt"
)

// 定义一个Animal类
type Animal struct {
	Color string
}

func (this *Animal) Eat(){
	fmt.Println("eating...")
}

func (this *Animal) walk(){
	fmt.Println("walking...")
}

type Dog struct {
	Animal  // 表示Dog类继承了Animal（结构体嵌套）
	Age int 
	Nick string
}

func (this *Dog) Eat(){  // 重写父类的方法
	fmt.Println("dog is eating ...")
}

func (this *Dog) Shout()  {  // 子类新增方法
	fmt.Println("Dog wang wang ...")
}

func main()  {
	// 创建父类对象
	a := Animal{Color: "red"}
	fmt.Println(a.Color)
	a.Eat()
	fmt.Println("========================")

	// 创建子类对象
	d := Dog{Animal{Color: "black"}, 18,  "XiaoHei"}
	d.Shout()
	fmt.Println("========================")
	// 或者
	var d1 Dog
	d1.Color = "black"
	d1.Age = 2
	d1.Nick = "XiaoHei"
	d1.Eat()
}





