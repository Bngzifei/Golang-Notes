package main

import "fmt"

/*
当通过通道发送有限的数据时,我们可以通过close函数关闭
通道来告知从该通道接收值的goroutine停止等待.
当通道被关闭时,往该通道发送值会引发panic,从该通道里接收
的值一直都是类型零值.那如何判断一个通道是否被关闭了呢?

我们来看下面这个例子:
*/

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// 开启goroutine从ch1中接收值,并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值 ok = false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range 循环
		fmt.Println(i)
	}
}

/*
从上面的例子中我们看到有两种方式在接收值的时候判断通道是否
被关闭,我们通常使用的是 for range 的方式
 */