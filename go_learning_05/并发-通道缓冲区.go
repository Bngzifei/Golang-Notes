package main

import "fmt"

/*
通道可以设置缓冲区,通过make的第二个参数指定缓冲区的大小:

ch := make(chan int, 100)

带缓冲区的通道允许发送端的数据发送和接收端的数据处于异步状态,
就是说发送端发送的数据可以放在缓冲区里面,可以等待接收端去取数据,
而不是立刻需要接收端去获取数据.

不过由于缓冲区的大小是有限的,所以还是必须有接收端来接收数据的,
否则缓冲区一满,数据发送端就无法再发送数据了.

注意:如果通道不带缓冲,发送方会阻塞直到接收方从通道中接收了值.如果通道
带缓冲,发送当则会阻塞直到发送的值被拷贝到缓冲区内;如果缓冲区已满,
则意味着需要等待直到某个接收方获取到一个值.接收方在有值可以接收之前
会一直阻塞.
 */

func main() {
	// 这里我们定义了一个可以存储整数类型的带缓冲通道
	// 缓冲区大小为2
	ch := make(chan int,2)

	// 因为 ch 是带缓冲的通道,我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2

	// 获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
输出结果: 1
          2
 */