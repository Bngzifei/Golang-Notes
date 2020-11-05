package main

import (
	"fmt"
	"time"
)

/*
时间到了,多次执行
*/

func main() {
	// 1.获取ticker对象
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	// 子协程
	go func() {
		for {
			// <- ticker.C
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				// 停止
				ticker.Stop()
			}
		}
	}()
	// 下面的这种 for{} 循环就是死循环 相当于 while True
	for {
	}
}
