// package main

// import "fmt"

// func main() {

// 	//仅用于接受数据
// 	mychnl1 := make(<-chan string)

// 	//仅用于发送数据
// 	mychnl2 := make(chan<- string)

// 	//显示通道的类型
// 	fmt.Printf("%T", mychnl1)
// 	fmt.Printf("\n%T", mychnl2)
// }

// 将双向通道转换为单向通道
package main

import "fmt"

func sending(s chan<- string) {
	s <- "nbooo"
}

func main() {

	//创建双向通道
	mychanl := make(chan string)

	//在这里，sending()函数将双向通道转换为仅发送通道
	go sending(mychanl)

	//在这里，通道只在goroutine内部发送，而在goroutine之外，通道是双向的，所以它打印nhooo
	fmt.Println(<-mychanl)
}
