package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a := 1
	b := 2
	c := a
	d := b
	a = d
	b = c
	fmt.Println(&a, b)

	// x := 1
	// y := 2
	// swap(&x, &y)
	// fmt.Println(x, y)
}
