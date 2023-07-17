package main

import (
	"fmt"
	"net/url"
	"time"
)

func main() {
	// go中的原时间格式time.Time类型
	NowTime := time.Now()
	fmt.Println("原时间格式time.Time类型:", NowTime)

	// Format意为格式， 将time.Time类型转为字符串（string）类型
	time1 := NowTime.Format("2006-01-02 15:04:05")
	fmt.Println("字符串(string)类型time2:", time1)

	// Parse意为解析， 将日期字符串格式转为时间
	layout := "2006-01-02 15:04:05"
	time2, err := time.Parse(layout, time1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("时间time2:", time2)

	// 时间戳
	now := time.Now().Unix()
	fmt.Println("时间戳now:", now)

	//把时间戳转为字符串（string）类型(int64 -> string)
	time3 := time.Unix(now, 0).Format("2006-01-02 15:04:05")
	fmt.Println("字符串(string)类型time3:", time3)

	beginOfMonth := time.Date(NowTime.Year(), NowTime.Month(), NowTime.Day(), 1, 0, 0, 0, NowTime.Location())
	fmt.Println("组装时间beginOfMonth:", beginOfMonth)
	endOfMonth := beginOfMonth.AddDate(0, 1, -1)
	fmt.Println("组装时间endOfMonth:", endOfMonth)

	t1 := time.Time{}
	fmt.Println("time.Time未初始化:", t1)
	t1Second := t1.Unix()
	fmt.Println("t1.Second:", t1Second)

	QureyTime()
}

// newbing
// 你可以使用 net/url 包中的 QueryUnescape 函数来解码 URL 编码的字符串。例如：
func QureyTime() {
	encodedTime := "2021-08-01%2000%3A00%3A00"
	decodedTime, err := url.QueryUnescape(encodedTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(decodedTime)
}
