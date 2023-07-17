// 恢复用于从紧急情况或错误情况中重新获得对程序的控制。
//它停止终止序列并恢复正常执行。从延迟函数中调用。它检索通过panic调用传递的错误值。通常，它返回nil，没有其他效果

package main

import "fmt"

func main() {
	fmt.Println(SaveDivide(10, 0))
	fmt.Println(SaveDivide(10, 10))
}
func SaveDivide(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	quotient := num1 / num2
	return quotient
}
