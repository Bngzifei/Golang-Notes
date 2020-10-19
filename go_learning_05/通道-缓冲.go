package main

import (
	"fmt"
	"time"
)

/*
通道带缓冲与不带缓冲的区别
*/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	fmt.Printf("sum:")
	fmt.Printf("%#v\n", sum)
	c <- sum // 把sum 发送到通道c
	fmt.Println("after channel pro")
}

// 通道不带缓冲,表示是同步的,只能向通道c发送一个数据,只要这个数据没被接收
// 然后所有的发送就被阻塞

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	fmt.Println("go [0,3]")
	// a
	go sum(s[:len(s)/2], c)

	// 这里开启一个新的运行期线程,这个是需要时间的,本程序继续往下走
	fmt.Println("go [3,6]")
	// b
	go sum(s[len(s)/2:], c)
	fmt.Println("go2 [0,3]")
	// c
	go sum(s[:len(s)/2], c)
	// d
	fmt.Println("go2 [3,6]")
	go sum(s[len(s)/2:], c)

	/*
		a b c d和main是一起争夺cpu的,他们的执行顺序完全无序,甚至里面不同的语句都相互
		穿插.但无缓冲的等待是同步的,所以接下来 a b c d都会执行,一直执行到c <- sum后,
		开始同步阻塞,因此,after channel pro 是打印不出来的,等要打印after channel pro的时候,
		main就结束了.
	*/
	fmt.Println("go3 start waiting...")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("go3 waited 1000 ms")

	// 因为a b c d都在管道门口等着,这里多一个,a b c d 就继续一个,这个结果的顺序可能是acbd
	aa := <-c
	bb := <-c

	fmt.Println(aa)
	fmt.Println(bb)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

}
