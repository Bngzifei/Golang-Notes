package main

import (
	"fmt"
	"time"
)

/*
go 语言支持并发,我们只需要通过go关键字来开启
goroutine即可.

goroutine是轻量级线程,goroutine的调度是由Golang
运行时进行管理的.

goroutine 语法格式:
 go 函数名(参数列表)

Go 允许使用go语句开启一个新的运行期线程,即goroutine,
以一个不同的,新创建的goroutine来执行一个函数.同一个程序
中的所有goroutine 共享同一个地址空间

这不就是协程么
*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
