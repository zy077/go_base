package main

import (
	"fmt"
)

func main(){
	var a string
	a = "haha"
	fmt.Println(a)

	var b interface{}
	b = a
	val, isStr := b.(string)
	if isStr{
		fmt.Printf("%s是string类型", val)
	}else {
		fmt.Printf("%s不是string类型", val)
	}
}
