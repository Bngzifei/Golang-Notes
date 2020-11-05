package main

import (
	"fmt"
	"sync"
)

/*
在代码中生硬的使用time.Sleep肯定是不合适的,go
语言中可以使用sync.WaitGroup来实现并发任务的同步.
sync.WaitGroup有以下几个方法:


sync.WaitGroup内部维护着一个计数器,计数器的值可以增加
和减少.例如当我们启动了N个并发任务时,就将计数器值增加N.
每个任务完成时通过调用Done()方法将计数器减1.通过调用
Wait()来等待并发任务执行完,当计数器值为0时,表示所有并发
任务已经完成.
*/

var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("Hello Goroutine")
}

func main() {
	wg.Add(1)
	// 启动另外一个goroutine去执行hello函数
	go hello()
	fmt.Println("main goroutine done!")
	wg.Wait()
}

