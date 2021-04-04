package main // 程序的包名

/*
导包
import "fmt"
import "time"
*/
import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Hello ")
	time.Sleep(1 * time.Second)
	fmt.Print("Go!")
}
