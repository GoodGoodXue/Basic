package main

import "fmt"

func main() {

	var i int = 0

	for i < 8 {
		if i == 5 {
			i = i + 2
			continue
		}
		fmt.Printf("i: %v\n", i)
		i++
	}

}
