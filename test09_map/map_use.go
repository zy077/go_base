package main

import (
	"fmt"
	"strings"
)

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

	// 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
	WordStr := "how do you do"
	CountWord(WordStr)

	// 打印
	PrintStr()
}

func forMap(myMay map[string]string)  {
	// map作为参数，传递的是引用（指针/内存地址值）
	for key, value := range myMay{
		fmt.Printf("%s = %s\n", key, value)
	}
	fmt.Println("===================")
}

//写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
func CountWord(WordStr string){
	Wordlist := make([]string, 0, 10)
	WordMap := make(map[string]int)
	Wordlist = strings.Split(WordStr, " ")
	for _, k := range Wordlist{
		if _, ok := WordMap[k]; ok {
			WordMap[k] += 1
		}else{
			WordMap[k] = 1
		}
	}
	for k, v := range WordMap {
		fmt.Printf("{%s: %d} \n", k, v)
	}
	fmt.Println("===================")
}

func PrintStr() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}