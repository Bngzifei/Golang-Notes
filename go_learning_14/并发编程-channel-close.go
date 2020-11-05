package main

import "fmt"

/*
可以通过内置的close()函数关闭channel(如果你的管道不往里存值
或者取值的时候,一定记得关闭管道)
*/

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main 结束")
}
