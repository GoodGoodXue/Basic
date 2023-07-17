package main

import (
	"fmt"
)

type users struct {
	id   int
	name string
	age  int
}

func main() {

	// var date []users
	users1 := users{
		id:   1,
		name: "wang",
		age:  18,
	}
	fmt.Println("person:", users1)

}
