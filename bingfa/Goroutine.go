package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}

func main() {

	//调用Goroutine
	go display("Welcome")

	//调用普通函数
	display("nhooo")

}
