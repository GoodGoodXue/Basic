// newbing
// 你可以使用 net/url 包中的 QueryUnescape 函数来解码 URL 编码的字符串。例如：
package pkg

import (
	"fmt"
	"net/url"
)

func QureyTime() {
	encodedTime := "2021-08-01%2000%3A00%3A00"
	decodedTime, err := url.QueryUnescape(encodedTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(decodedTime)
}
