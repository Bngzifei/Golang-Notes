// Go语言中,一个break的作用范围是该语句出现后的最内部的结构,它可以被作用于任何形式的for循环(计数器,条件判断等).但在switch或select语句中,break语句的作用结果是跳过整个代码块,执行后续的代码.

// 循环嵌套循环时,可以在break后指定标签,用标签决定哪个循环被终止

// 跳出指定循环:
package main

import (
	"fmt"
)

func main() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}
}
