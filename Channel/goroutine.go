package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	wg.Done()
}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go download("a.com/" + string(i+'0'))
		// 这段代码中，string(i+'0') 的作用是将整数 i 转换为其对应的字符串形式。在 Go 语言中，字符和整数可以相互转换。例如，字符 '0' 的整数值为 48，因此 '0' + 1 的结果为 49，即字符 '1'。所以，在这段代码中，string(i+'0') 将整数 i 转换为其对应的数字字符。

		// 在 Go 语言中，如果你尝试使用 string(i) 将整数转换为字符串，你会得到一个表示该整数对应 Unicode 码点的字符。例如，string(48) 的结果是 "0"，而不是 "48"。

		// 如果你想将整数转换为其对应的字符串形式，可以使用 strconv.Itoa() 函数。例如，strconv.Itoa(48) 的结果是 "48"。
	}
	wg.Wait()
	fmt.Println("Done!")
}
