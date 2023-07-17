//vote 实例
// // 这里的val是字符串变量（string）
// val, err := items.Client.Get(ctx, key_vote).Result()
// if err != nil {
// 	panic(err)
// }
// // var p2 []CmsV2VoteItems
// // json.Unmarshal() 函数用于将 JSON 字符串解码为 Go 语言中的数据结构。
// // 它的第一个参数是一个字节切片，包含要解码的 JSON 数据；第二个参数是一个指向要存储解码结果的变量的指针。
// // 将 val 中的 JSON 字符串解码为(存入) Voteinfo 变量
// json.Unmarshal([]byte(val), &Voteinfo)

// json.Unmarshal() 函数用于将 JSON 字符串解码为 Go 语言中的数据结构。
// 它的第一个参数是一个字节切片，包含要解码的 JSON 数据； 第二个参数是一个指向要存储解码结果的变量的指针。

// 例如，假设我们有一个 JSON 字符串表示一个人的信息，如下所示：
// personJSON := `{"name": "Alice", "age": 30}`

// 我们可以定义一个 Person 类型来表示这个人的信息，并使用 json.Unmarshal() 函数将 JSON 字符串解码为这个类型的变量：

package main

import "encoding/json"

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var p Person
	err := json.Unmarshal([]byte(personJSON), &p)
	if err != nil {
		// 处理错误
	}
}

// 在上面的代码中，我们定义了一个 Person 类型，其中包含两个字段：Name 和 Age。
// 然后，我们创建了一个 Person 类型的变量 p，并使用 json.Unmarshal() 函数将 JSON 字符串解码为这个变量。
