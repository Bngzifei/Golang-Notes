package main

import "fmt"

/*
select是Go中的一个控制结构,类似于用于通信的switch语句.
每个case必须是一个通信操作,要么是发送,要么是接收.

select随机执行一个可运行的case.如果没有case可运行,它将阻塞,
直到有case可运行.一个默认的子句应该总是可运行的.

以下描述了select语句的语法:
	每个case都必须是一个通信
	所有channel表达式都会被求值
	所有被发送的表达式都会被求值
	如果任意某个通信可以进行,它就执行,其他被忽略
	如果有多个case都可以运行,select会随机公平地选出一个执行.其他不会执行
	否则:
		1.如果有default子句,则执行该语句
		2.如果没有default子句,select将阻塞,直到某个通信可以运行;
		go不会重新对channel或值进行求值

golang 的 操作符号"<-"  到底是代表什么？
	<-是对chan类型来说的。chan类型类似于一个数组。
	当 <- chan 的时候是对chan中的数据读取；
	相反 chan <- value 是对chan赋值。

帖子链接
https://bbs.csdn.net/topics/391922084

*/
func main() {
	var c1, c2, c3 chan int
	var i1, i2 int

	select {
		// 从chan中读取
		case i1 = <-c1:
			fmt.Printf("received ", i1, " from c1\n")
		// 对chan进行赋值
		case c2 <- i2:
			fmt.Printf("sent ", i2, " to c2\n")
		case i3, ok := (<-c3):
			if ok {
				fmt.Printf("received ", i3, " from c3\n")
			} else {
				fmt.Printf("c3 is closed\n")
			}
		default:
			fmt.Printf("no communication")
	}
}
