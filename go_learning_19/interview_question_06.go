package main

import "fmt"

// 基于类型int创建了新类型MyInt1
type MyInt1 int
// 创建了int的类型别名MyInt2.
type MyInt2 = int

// cannot use i (type int) as type MyInt1 in assignment
func main() {
	var i int = 10
	// 类型别名与类型定义的区别
	// go是强类型语言,将 int 类型的变量赋值给 MyInt1 类型的变量
	var i1 MyInt1 = i
	// 使用强制类型转化:
	var i1 MyInt1 = MyInt1(i)
	// MyInt2 只是 int 的别名，本质上还是 int，可以赋值
	var i2 MyInt2 = i
	fmt.Println(i1,i2)
}
