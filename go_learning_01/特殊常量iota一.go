package main

import "fmt"

/*
iota,特殊常量,可以认为是一个 可以被编译器修改的常量
iota 在const关键字出现时将被重置为0(const 内部的第一行之前),
const中每新增一行常量声明将使iota计数一次,
(iota 可理解为 const 语句块中的行索引)
iota可以被用作枚举值:

const (
    a = iota
    b = iota
    c = iota
)

第一个iota等于0,每当iota在新的一行被使用时,它的值都会自动加1,
所以 a = 0, b = 1, c = 2 可以简写为如下形式:
const (
    a = iota
    b
    c
)

*/

func main() {
	const (
		a = iota // 0
		b        // 1
		c        // 2
		d = "ha" // 独立值，iota += 1
		e        // "ha"   iota += 1
		f = 100  // 100  iota +=1
		g        // 100  iota +=1
		h = iota // 7, 恢复计数
		i        // 8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
