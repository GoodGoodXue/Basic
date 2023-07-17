package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	var val uint64 = 730
	var ratio uint64 = 50000
	res := val * ratio / 1e6
	fmt.Println("res:", res)

	var val2 float64 = 730
	var ratio2 float64 = 50000
	res2 := val2 * (ratio2 / 1000000.00)
	fmt.Println("res2:", res2)
}
