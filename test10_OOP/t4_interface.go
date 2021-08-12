package main

import (
	"fmt"
)

// 父类（接口），本质是一个指针
type AnimalIF interface {
	// 接口就是抽象方法的集合
	// 声明抽象方法
	 Sleep()
	GetColor() string
	GetType() string
}

// 声明一个Dog子类
type Pig struct {
	// 继承Animal类（可省略不写）
	Color string
}

// 实现父类的所有抽象方法
func (this *Pig) Sleep()  {
	fmt.Println("Pig is sleeping ...")
}

func (this *Pig) GetColor() string {
	return this.Color
}

func (this *Pig) GetType() string {
	return "Dog"
}

// 声明一个Cat子类
type Cat struct {
	// 继承Animal类（可省略不写）
	Color string
}

// 实现父类的所有抽象方法
func (this *Cat) Sleep()  {
	fmt.Println("Dog is sleeping ...")
}

func (this *Cat) GetColor() string {
	return this.Color
}

func (this *Cat) GetType() string {
	return "Cat"
}

func Show(animal AnimalIF)  {  // 接口本质是一个指针
	animal.Sleep()
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())
	fmt.Println("========================")
}

// main方法
func main()  {
	// 声明一个父类类型对象
	var animal AnimalIF
	animal = &Pig{Color: "Black"}   // 接口本质是指针，故需要传递地址值
	animal.Sleep()
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())
	fmt.Println("========================")
	animal = &Cat{Color: "yellow"}
	animal.Sleep()
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())
	fmt.Println("========================")

	// 或者
	Show(&Cat{Color: "yellow"})  // 需要传递引用
	Show(&Pig{Color: "black"})
}