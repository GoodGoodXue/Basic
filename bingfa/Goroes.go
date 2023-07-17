// 多个Goroutine
package main

import (
	"fmt"
	"time"
)

// goroutine1
func Aname() {

	arr1 := [4]string{"Rohint", "Suman", "Aman", "Ria"}

	for t1 := 0; t1 <= 3; t1++ {

		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%vs\n", arr1[t1])
	}
}

// goroutine 2
func Aid() {
	arr2 := [4]int{300, 301, 302, 303}

	for t2 := 0; t2 <= 3; t2++ {

		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%d\n", arr2[t2])
	}
}

func main() {

	fmt.Println("!...主 Go-routine 开始...!")

	//调用 Goroutine 1
	go Aname()

	//调用 Goroutine 2
	go Aid()

	time.Sleep(3500 * time.Millisecond)
	fmt.Println("\n!...主 Go-routine 结束...!")
}
