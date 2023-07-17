package main

import "fmt"

func modify(Z int) {
	Z = 70
}

func main() {

	var Z int = 10

	fmt.Println("函数调用前,Z的值为 =", Z)

	//按值调用
	modify(Z)

	fmt.Println("\n函数调用后, Z的值为 =", Z)
}
