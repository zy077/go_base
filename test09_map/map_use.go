package main

import "fmt"

func main()  {
	// 声明一个map
	map0 := make(map[string]string)
	// 添加数据
	map0["陕西"] = "西安"
	map0["湖南"] = "长沙"
	map0["江西"] = "南昌"

	// 遍历
	for key, value := range map0{
		fmt.Printf("%s = %s\n", key, value)
	}
	fmt.Println("===================")

	// 修改
	map0["陕西"] = "长安"

	// 函数遍历
	forMap(map0)

	// 删除
	delete(map0, "陕西")
	forMap(map0)

}

func forMap(myMay map[string]string)  {
	// map作为参数，传递的是引用（指针/内存地址值）
	for key, value := range myMay{
		fmt.Printf("%s = %s\n", key, value)
	}
	fmt.Println("===================")
}