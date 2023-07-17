package main

import "fmt"

func area(length, whidth int) int {

	Ar := length * whidth
	return Ar

}

func main() {

	fmt.Printf("矩形的面积 : %d", area(12, 10))
}

// package main

// import "fmt"

// //area()用于查找矩形的面积
// //函数的两个参数，即长度和宽度
// func area(length, width int) int {

// 	Ar := length * width
// 	return Ar
// }

// func main() {

// 	//显示矩形的面积
// 	//通过方法调用
// 	fmt.Printf("矩形的面积 : %d", area(12, 10))
// }
