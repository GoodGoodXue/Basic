package main

import "fmt"

func main() {

	//创建通道
	c := make(chan int)
	select {
	case <-c:
	default:
		fmt.Println("!..Default case..!")
	}
}
