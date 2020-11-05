package main

import "fmt"

/*
有的时候我们会将通道作为参数在多个任务函数间传递,
很多时候我们在不同的任务函数中使用通道都会对其进行限制,
比如限制通道在函数中只能发送或只能接收.

go 语言中提供了单向通道来处理这种情况.例如,我们把上面的
例子改造如下:
*/

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		// 存到通道中
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	// 对 in 通道进行遍历取值,然后存到 out 通道
	for i := range in {
		// 存到通道中
		out <- i * i
	}
	close(out)
}


func printer(in <-chan int) {
	// 遍历 in 通道的值
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go counter(ch1)
	go squarer(ch2, ch1)
	//printer(ch1)
	//fmt.Printf("------------------")
	printer(ch2)
}

/*
其中,
1. chan <- int 是一个只能发送的通道,可以发送但是不能接收
2. <- chan int 是一个只能接收的通道,可以接收但是不能发送

在函数传参及任何赋值操作中将双向通道转换为单向通道是可以的,但反过
来时不可以的.

注意:
关闭已经关闭的channel也会引发panic
 */