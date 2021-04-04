package main

import (
    "encoding/json"
    "fmt"
)

type Movie struct{
    Title string    `json:"电影名字"`  // 该属性在json格式的字符串中显示"电影名字"
    TicketPrice int `json:"票价"`
    Actors []string `json:"演员表"`
}


func main(){
    // 声明一个对象
    movie := Movie{"泰坦尼克号", 20, []string{"Jack", "Rose"}}
    // 编码 结构体——>json (序列化)
    jsonStr, err := json.Marshal(movie)
    if err != nil {
        fmt.Println("序列化出错！")
        return
    }
    fmt.Printf("jsonStr = %s\n", jsonStr)

    // 解码 json——>结构体 （反序列化）
    myMovie := Movie{}
    err = json.Unmarshal(jsonStr, &myMovie)
    if err != nil {
        fmt.Println("反序列化错误！")
    }
    fmt.Printf("myMovie = %s\n", myMovie)
}