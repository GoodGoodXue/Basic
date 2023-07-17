package main

import "fmt"

func main() {

	c := make(chan int)
	//var c chan int
	select {
	case <-c:
	//case x1 <-c:
	// fmt.Println("Value: ", x1)
	//出现思索，执行默认语句避免死锁
	default:
		fmt.Println("!..Default case..!")

	}
}
