package main

import "fmt"

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

func main() {
	defer_call()
}

/*
打印后
打印中
打印前
panic: 触发异常

defer的执行顺序是后进先出.当出现panic语句的时候,会先按照defer的
后进先出的顺序执行,最后才会执行panic
 */
