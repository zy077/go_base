package main

import "fmt"

func main() {
	// 声明map有三种方式
	// 方式一:
	var map01 map[string]string  // 声明一个map01是一种map类型，key为string类型，value也是string类型，默认为map为空
	if map01 == nil {
		fmt.Println("map01 为空map")
	}
	map01 = make(map[string]string, 10)  // 给map01开辟容量为10的空间（扩容机制和array相同）
	map01["name"] = "TongTianZi"  // 给map01添加元素
	map01["age"] = "18"
	map01["gender"] = "women"
	fmt.Println(map01)
	fmt.Println("====================")

	// 方式二：不指定容量，表示容量为任意个
	map02 := make(map[string]string)
	map02["first"] = "python"
	map02["second"] = "go"
	map02["third"] = "java"
	fmt.Println(map02)
	fmt.Println("====================")

	// 方式三：声明一个map，并进行初始化
	map03 := map[int]string{
		1: "python",
		2: "go",
		3: "java",
	}
	fmt.Println(map03)
}
