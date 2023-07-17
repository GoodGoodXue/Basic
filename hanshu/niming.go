// package main

// import "fmt"

// func main() {

// 	//匿名函数
// 	func() {
// 		fmt.Println("Welcome! to (cainiaojc.com)")
// 	}()
// }

// package main

// import "fmt"

// func main() {

// 	//分配一个匿名函数到一个变量
// 	value := func() {
// 		fmt.Println("Wecome! to (cainiaojc.com)")
// 	}
// 	value()
// }

package main

import "fmt"

func main() {

	//在匿名函数传递参数
	func(ele string) {
		fmt.Println(ele)
	}("nhooo")

}
