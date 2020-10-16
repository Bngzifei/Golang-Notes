package main

import "fmt"

/*
在函数体外声明的变量称之为全局变量,全局变量可以在整个包甚至
外部包(被导出后)使用.
全局变量可以在任何函数中使用,以下实例演示了如何使用全局变量:
*/

// 声明全局变量
var g int

func main() {
	var a, b int

	// 初始化参数
	a = 10
	b = 20
	g = a + b

	fmt.Printf("结果: a = %d, b = %d and g = %d\n", a, b, g)

}

/*
go 语言程序中全局变量与局部变量名称可以相同,但是函数内的局部变量会被
优先考虑.实例如下:
 */


///* 声明全局变量 */
//var g int = 20
//
//func main() {
//	/* 声明局部变量 */
//	var g int = 10
//	// 实际输出的是局部变量 10
//	fmt.Printf ("结果： g = %d\n",  g)
//}
