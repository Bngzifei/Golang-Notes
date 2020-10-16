package main

import "fmt"

func main() {
	// 定义局部变量
	var a int = 10

	for a < 20 {
		fmt.Printf("a 的值为 : %d\n", a)
		a++
		if a == 15 {
			// 跳出循环体,执行后续语句
			break
		}
	}
	fmt.Printf("跳出来了...")

}
