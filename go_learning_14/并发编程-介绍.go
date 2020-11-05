package main

import (
	"fmt"
	"sync"
)

/*
goroutine奉行通过通信来共享内存,而不是共享内存来通信.

goroutine 的概念类似于线程,但goroutine是由go的运行时
调度和管理的.go程序会智能的将goroutine中的任务合理地分配
给每个CPU.go语言之所以被称之为现代化的编程语言,就是因为它
在语言层面已经内置了调度和上下文切换的机制.

在go语言编程中你不需要去自己写进程,线程和协程,你的技能包
里只有一个技能-goroutine.当你需要让某个任务并发执行的时候,
你只需要把这个任务包装成一个函数,开启一个goroutine去执行
这个函数就可以了.就是这么简单粗暴.

go语言中使用goroutine非常简单,只需要在调用函数的时候在前面加上
go关键字,就可以为一个函数创建一个goroutine.

一个goroutine必定对应一个函数,可以创建多个goroutine去执行
相同的函数.

启动单个goroutine
启动goroutine的方式非常简单,只需要在调用的函数(普通函数和匿名函数)
前面加上一个go关键字

举个例子如下:
*/

var wg sync.WaitGroup // 使用sync.WaitGroup来实现goroutine的同步

func hello(i int) {
	defer wg.Done() // goroutine结束就登记 -1
	fmt.Println("Hello Goroutine!", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记 +1
		go hello(i)
	}

	wg.Wait() // 等待所有登记的goroutine都结束
}


/*
多次执行上面的代码,会发现每次打印的数字的顺序都不一致,这是因为10个goroutine
是并发执行的,而goroutine的调度是随机的.

注意:如果主协程退出了,其他任务还执行吗?
 */