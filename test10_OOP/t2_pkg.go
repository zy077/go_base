package main

import (
	"fmt"
)

// 定义一个Person的结构体
//如果类名首字母大写，表示其他包也能够访问
type Person struct {
	//如果说类的属性首字母大写, 表示该属性是对外能够访问的，否则的话只能够类的内部访问
	Name   string
	age    int
	Gender string
}

// 给结构体绑定方法

/*
func (this Person) Show() {
	fmt.Printf("name = %s, age = %d, gender = %s\n", this.Name, this.age, this.Gender)
	fmt.Println("==================")
}

func (this Person) GetInfo() (string, int, string) {
	return this.Name, this.age, this.Gender
}

func (this Person) SetInfo(name string, age int, gender string) {
	this.Name = name
	this.age = age
	this.Gender = gender
}
*/

// 方法名首字母大写表示对外开放（即其他包中也能调用），否则仅限于本包中使用
func (this *Person) Show() {
	// this表示调用该方法的对象的副本（拷贝）
	fmt.Printf("name = %s, age = %d, gender = %s\n", this.Name, this.age, this.Gender)
	fmt.Println("==================")
}

func (this *Person) GetInfo() (string, int, string) {
	return this.Name, this.age, this.Gender
}

func (this *Person) SetInfo(name string, age int, gender string) {
	this.Name = name
	this.age = age
	this.Gender = gender
}

func main() {
	// 创建一个对象
	p := Person{Name: "TongTianZi", age: 18, Gender: "women"}
	p.Show()
	// 修改属性
	p.SetInfo("高圆圆", 20, "女")
	p.Show()
}
