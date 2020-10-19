package main

import "fmt"

/*
go通过range关键字来实现遍历读取到的数据,类似于数组或切片.
格式如下:  v,ok := <-ch

如果通道接收不到数据后ok就为false,这时通道就可以使用close函数
来关闭。
*/

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据,因为c在发送完10个数据之后就关闭了
	// 通道.所以这里我们range函数在接收到10个数据之后就结束了.
	// 如果上面的c通道不关闭,那么range函数就不会结束,从而在接收第11个数据的
	//时候就阻塞了.
	for i := range c {
		fmt.Println(i)
	}
}

/*
goroutine 是golang中在语言级别实现的轻量级线程,仅仅利用go就能立刻起一个新线程.
多线程会引入线程之间的同步问题,在golang中可以使用channel作为同步的工具.

通过channel可以实现两个goroutine之间的通信.

创建一个channel,make(chan TYPE{,NUM})
TYPE指的是channel中传输的数据类型,第二个参数是可选的,指的是channel容量大小

向channel传入数据,CHAN <- DATA, CHAN指的是目的channel即收集数据的一方.
DATA则是要传的数据.

从channel读取数据,DATA := <- CHAN, 和向channel传入数据相反,在数据输送箭头的
右侧是channel,形象地展现了数据从隧道流出到变量里

关闭通道并不会丢失里面的数据,只是让读取通道数据的时候不会读完之后一直阻塞等待新
数据写入


无缓冲和有缓冲的区别:
无缓冲是同步的,例如make(chan int),就是一个送信人去你家门口送信,你不在家他不走,
你一定要接下信,他才会走,无缓冲保证信能到你手上.

有缓冲是异步的,例如make(chan int, 1) ,就是一个送信人去你家扔到你家的信箱,转身
就走,除非你的信箱满了,他必须等信箱空下来,有缓冲的保证信能进你家的邮箱.

通道遵循先进先出原则.
不带缓冲区的通道在向缓冲区发送值时,必须及时接收,且必须一次接收完成.
而带缓冲区的通道则会以缓冲区满而阻塞,直到先塞发送到通道的值被从通道中
接收才可以继续往通道传值.就像往水管里推小钢珠一样,如果钢珠塞满没有从另一头放出,那么这
一头就没法再往里塞,是一个道理.
 */