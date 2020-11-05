package main

import (
	"fmt"
	"time"
)

/*
在某些场景下我们需要同时从多个通道接收数据.
通道在接收数据时,如果没有数据可以接收将会发生
阻塞.

你也许会写出如下代码使用遍历的方式来实现:
for{
    // 尝试从ch1接收值
    data, ok := <-ch1
    // 尝试从ch2接收值
    data, ok := <-ch2
    …
}

这种方式虽然可以实现从多个通道接收值的需求,但是运行性能
会差很多.为了应对这种场景,go内置了select关键字,可以同时
响应多个通道的操作

select的使用类似switch语句,它有一系列case分支和一个默认的分支,
每个case会对应一个通道的通信(接收或发送)过程.select会一直
等待,直到某个case的通信操作完成时,就会执行case分支对应的语句.
具体格式如下:
	select {
		case <- chan1:
			// 如果chan1 成功读取到数据,则进行该case处理语句
		case chan2 <- 1:
			// 如果成功向chan2写入数据,则进行该case处理语句
		default:
			// 如果上面都没有成功,则进入default处理流程
	}

select可以同时监听一个或多个channel,直到其中一个channel ready

*/

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
// 如果多个channel同时ready,则随机选择一个执行
// 设置两个的休眠时间一样,可以看到随机选择一个执行
func test2(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test2"
}

func main() {
	// 2个管道
	output1 := make(chan string)
	output2 := make(chan string)

	// 跑2个子协程,写数据
	go test1(output1)
	go test2(output2)
	// 用select监控,实际效果就是看 哪个通道先准备好
	// 上面的test1休眠5秒后准备好,所以下面的select执行 s1
	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}
}
