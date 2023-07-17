package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mu sync.Mutex

func main() {
	for i := 0; i < 1000; i++ {
		// func(){} 为匿名函数 // go ()调用该匿名函数
		go func() {
			// 锁上，每一个只能一个goroutine
			mu.Lock()
			count = count + 1
			// 计算完解锁，下一个goroutine
			mu.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(count)
}
