package main

import (
    "fmt"
    "strings"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
    coins = 50
    users = []string{
        "Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
    }
    DistributionMap = make(map[string]int, len(users))
)

func dispatchCoin(str string) int {
    sum := 0
    str_list := strings.Split(str, "")
    for _, val := range str_list {
        if val == "e" || val == "E" {
            sum += 1
            coins -= 1
        }
        if val == "i" || val == "I" {
            sum += 2
            coins -= 2
        }
        if val == "o" || val == "O" {
            sum += 3
            coins -= 3
        }
        if val == "u" || val == "U" {
            sum += 4
            coins -= 4
        }
    }
    return sum
}

func main() {
    for _, str := range users {
       DistributionMap[str] = dispatchCoin(str)
    }
    fmt.Println(DistributionMap)
    fmt.Printf("剩余金币：%d \n", coins)
}
