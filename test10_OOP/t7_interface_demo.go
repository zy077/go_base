package main

import "fmt"

type Mover interface {
    move(name string)
}

type Car struct {
    ID int `json:"id"`
    Name string `json:"name"`
}

type Bus struct {
    ID int `json:"id"`
    Name string `json:"name"`
}

// 接受者是struct类型, 调用方可以是struct，也可以是*struct
func (c Car) move(name string) {
    fmt.Printf("%s is running ... \n", name)
}

// 接受者是struct类型的指针，调用者只能是*struct
func (b *Bus) move (name string) {
    fmt.Printf("%s is running ... \n", name)
}

func main() {
    var m Mover
    m = Car{}
    m.move("比亚迪")
    m = &Car{}
    m.move("大众")
    m = &Bus{}
    m.move("客车")
}
