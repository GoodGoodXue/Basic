// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	GoA()
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("main")
// }

// func GoA() {
// 	defer (func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("panic:" + fmt.Sprintf("%s", err))
// 		}
// 	})()

// 	go GoB()
// }

// func GoB() {
// 	panic("error")
// }

package main

import (
	"fmt"
)

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// defer为函数返回前的执行，多个defer则进行先进后出

// 1, 2
// (B, 1, 2) B, 1, 2, 3, ret=3

// defer1 返回前第二执行 带入值 A，1，3

// tmp2 = (D, 3, 2) D, 3, 2, 5, ret=5

// defer2 返回前第一执行 带入值 C，3，5
// C, 3, 0, 5, ret=3

// defer2: A, 1，3，4，ret=4

// defer无法修改匿名返回值
func incr(a int) int {
	var b int
	defer func() {
		a++
		b++
	}()
	a++
	b = a
	return b
}

// defer 可以修改【具名返回值】
func incr1(a int) (b int) {
	defer func() {
		a++
		b++
	}()
	a++
	b = a
	// return可加可不加 b
	return
}

func main() {

	// defer 执行顺序
	// x := 1
	// y := 2
	// tmp1 := calc("B", x, y)

	// defer 是带入上面值，再进行待执行的
	// defer calc("A", x, tmp1)
	// x = 3
	// tmp2 := calc("D", x, y)
	// defer calc("C", x, tmp2)
	// y = 4

	// defer 是否修改返回值
	var a, b int
	b = incr(a)
	fmt.Println(a, b)
}
