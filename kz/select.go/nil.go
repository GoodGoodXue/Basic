package main

import "fmt"

func main() {

	//创建通道
	var c chan int

	select {
	case x1 := <-c:
		fmt.Println("Value:", x1)
	default:
		fmt.Println("Default case..!")
	}
}
