// package main

// import "fmt"

// func main() {
// 	fmt.Println("c")
// 	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
// 		fmt.Println("d")
// 		if err := recover(); err != nil {
// 			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
// 		}
// 		fmt.Println("e")
// 	}()
// 	f()
// 	fmt.Println("f") //这里开始下面代码不会再执行
// }

// func f() {
// 	fmt.Println("a")
// 	panic("异常信息")
// 	fmt.Println("b") //这里开始下面代码不会再执行
// 	fmt.Println("f")
// }

package main

import "fmt"

func main() {
	fmt.Println("main start")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recober1", err.(string))
		}
	}()

	func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recover2", err.(string))
			}
		}()
		panic("panic exeption2!")
	}()
	panic("panic exception1!")
	fmt.Println("main end")
}

//  main函数被执行，压入运行时栈中
//  main函数开头的输出main start
//  接下来匿名defer被压入”栈1“中
//  执行匿名函数func(){}()，匿名函数被压入运行时栈中
//  匿名函数中defer被压入”栈2“中
//  抛出panic()
//  匿名函数中的defer被调用，捕获panic，并输出recover2 panic exception2!
//  匿名函数调用结束，匿名函数出运行时栈
//  抛出panic()
//  main函数中的defer被调用，不会panic，并输出recover1 panic exception1!
//  main函数出运行时栈
//  程序结束
