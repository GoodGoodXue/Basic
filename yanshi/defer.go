// package main

// import "fmt"

// func main() {
// 	arr := [5]int{}
// 	for i := range arr {
// 		defer fmt.Println(i)
// 	}
// }

// 闭包
package main

import "fmt"

func main() {
	arr := [5]int{}
	for i := range arr {
		defer func() { fmt.Println(i) }()
	}
}
