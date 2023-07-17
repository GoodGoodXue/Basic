package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome!! to Main function")

	//创建匿名Goroutine
	go func() {
		fmt.Println("Welcome!! to (cainiaojc.com)")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("GoodBye!! to Main function")
}
