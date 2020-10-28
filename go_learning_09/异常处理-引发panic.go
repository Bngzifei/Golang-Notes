package main

import "fmt"

/*
向已关闭的通道发送数据会引发panic
 */

func main(){
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var ch chan int = make(chan int,10)
	close(ch)
	ch <- 1

}

/*
send on closed channel
 */