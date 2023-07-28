package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // 解析参数

	if err != nil {
		return
	}

	fmt.Println(r.Form) //
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fprintf, err := fmt.Fprintf(w, "Hello astaxie!") // 写入到w的是输出到客户端
	if err != nil {
		return
	}

	fmt.Println(fprintf)
}

func main() {
	http.HandleFunc("/", sayhelloName)       // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
