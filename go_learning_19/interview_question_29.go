package main

import (
	"fmt"
	"time"
)

//func main() {
//	v := []int{1, 2, 3}
//	fmt.Println(v)
//	for i, value := range v {
//		// i -> index
//		// value -> element
//		v = append(v, value)
//		fmt.Println(i)
//	}
//	fmt.Println(v)
//}
// 不会出现死循环,能正常结束
// 循环次数在循环开始前就已经确定,循环内改变切片的长度,不影响循环次数.

func main() {
	var m = [...]int{1, 2, 3}
	// i : 0, 1, 2
	for i, v := range m {
		// 这里的 := 会重新声明变量,而不是重用
		//i := i
		//v := v
		go func(i, v int) {
			fmt.Println(i, v)
		}(i, v)
	}
	time.Sleep(time.Second * 3)
}

// for range 使用短变量声明(:=)的形式迭代变量,需要注意的是,变量i,v在每次
// 循环体中都会被重用,而不是重新声明.
// 各个goroutine中输出的i,v值都是for range循环结束后的i,v最终值,而不是各个
// goroutine启动时的i,v值.可以理解为闭包引用,使用的上下文环境的值.
// https://tonybai.com/2015/09/17/7-things-you-may-not-pay-attation-to-in-go/
