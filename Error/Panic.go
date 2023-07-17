// Panic是一种我们用来处理错误情况的机制。
//紧急情况可用于中止函数执行。当一个函数调用panic时，它的执行停止，并且控制流程到相关的延迟函数。

package main

import "os"

func main() {
	panic("Error Situation")
	_, err := os.Open("/tmp/file")
	if err != nil {
		panic(err)
	}
}
