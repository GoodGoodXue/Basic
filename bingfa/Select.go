package main

import (
	"fmt"
	"time"
)

// 函数1
func portal1(channel1 chan string) {

	time.Sleep(3 * time.Second)
	channel1 <- "Welcome to channel 1"
}

// 函数2
func portal2(channel2 chan string) {

	time.Sleep(9 * time.Second)
	channel2 <- "Welcome to channel 2"
}

func main() {

	//创建通道
	R1 := make(chan string)
	R2 := make(chan string)

	//使用goroutine调用函数1和函数2
	go portal1(R1)
	go portal2(R2)

	select {

	//case 1
	case op1 := <-R1:
		fmt.Println(op1)

		// case 2
	case op2 := <-R2:
		fmt.Println(op2)

	}

}
