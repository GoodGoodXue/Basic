// 创建通道
// package main

// import "fmt"

// func main() {

// 	//使用var关键字创建通道
// 	var mychannel chan int
// 	fmt.Println("channel的值:", mychannel)
// 	fmt.Printf("channel的类型: %T ", mychannel)

// 	//使用 make() 函数创建通道
// 	mychannel1 := make(chan int)
// 	fmt.Println("channel1的值: ", mychannel1)
// 	fmt.Printf("channel1的类型: %T ", mychannel1)
// }

// 从通道发送和接收数据
// package main

// import "fmt"

// func myfunc(ch chan int) {
// 	fmt.Println(234 + <-ch)
// }

// func main() {
// 	fmt.Println("主方法开始")
// 	//创建通道1
// 	ch := make(chan int)

// 	go myfunc(ch)
// 	ch <- 23
// 	fmt.Println("主方法结束")

// }

// 关闭通道
// 您也可以在close()函数的帮助下关闭通道。这是一个内置函数，并设置一个标识，表示不再有任何值将发送到该通道
// package main

// import "fmt"

// func myfun(mychnl chan string) {

// 	for v := 0; v < 4; v++ {
// 		mychnl <- "nhooo"
// 	}
// 	close(mychnl)
// }

// func main() {

// 	//  创建通道
// 	c := make(chan string)

// 	//使用 Goroutine
// 	go myfun(c)

// 	//当ok的值为true时，表示通道已打开，可以发送或接受数据
// 	//当ok的值为false时，表示通道已关闭
// 	for {
// 		res, ok := <-c
// 		if ok == false {
// 			fmt.Println("通道关闭 ", ok)
// 			break
// 		}
// 		fmt.Println("通道打开 ", res, ok)
// 	}
// }

// 重要事项
// package main

// import "fmt"

// func main() {

// 	// 使用 make() 函数创建通道
// 	mychnl := make(chan string)

// 	// 匿名 goroutine
// 	go func() {
// 		mychnl <- "GFG"
// 		mychnl <- "gfg"
// 		mychnl <- "Geeks"
// 		mychnl <- "nhooo"
// 		close(mychnl)
// 	}()

// 	//使用for循环
// 	for res := range mychnl {
// 		fmt.Println(res)
// 	}
// }

// 通道长度：在通道中，您可以使用len()函数找到通道的长度。在此，长度表示在通道缓冲区中排队的值的数量
// package main

// import "fmt"

// func main() {

// 	//使用 make() 函数创建通道
// 	mychnl := make(chan string, 4)
// 	mychnl <- "GFG"
// 	mychnl <- "gfg"
// 	mychnl <- "Geeks"
// 	mychnl <- "nbooo"

// 	// 使用 len() 函数查找通道的长度
// 	fmt.Println("channel长度为: ", len(mychnl))
// }

// 通道的容量：在通道中，您可以使用cap()函数找到通道的容量。在此，容量表示缓冲区的大小
package main

import "fmt"

func main() {

	// 使用 make() 函数创建通道
	mychnl := make(chan string, 5)
	mychnl <- "GFG"
	mychnl <- "gfg"
	mychnl <- "Geeks"
	mychnl <- "nbooo"

	//使用 cap() 函数查找通道的容量
	fmt.Println("channel容量为: ", cap(mychnl))
}
