package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    // 标签就是对结构体属性的补充说明
    Name string `info:"姓名" nick:"Yang"`
    Age int `info:"year"`
}

func FindTag(args interface{}){
    els := reflect.TypeOf(args).Elem()
    for i := 0; i < els.NumField(); i++ {
        j := els.Field(i).Tag.Get("info")
        k := els.Field(i).Tag.Get("nick")
        fmt.Println("info = ", j, "nick = ", k)
    }
}


func main(){
    var p Person
    FindTag(&p)
}
