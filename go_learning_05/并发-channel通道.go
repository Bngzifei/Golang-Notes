package main

import "fmt"

/*
通道(channel是用来传递数据的一个数据结构
通道可用于两个goroutine之间通过传递一个指定类型的值
来 同步运行和通讯.
操作符 <- 用于指定通道的方向,发送或接收.
如果未指定方向,则为双向通道.

ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据
           // 并把值赋给 v

声明一个通道很简单,我们是有chan关键字即可,通道在使用前必须先创建:
ch := make(chan int)

注意:默认情况下,通道是不带缓冲区的.发送端发送数据,同时必须有接收端
相应的接收数据(这意思就是 有发送端 就必须有接收端 )

下面的实例通过两个goroutine来计算数字之和,在goroutine完成计算之后,
它会计算两个结果的和:
*/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把sum发送到通道c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// 从通道c中接收
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
