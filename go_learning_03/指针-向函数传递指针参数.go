package main

import "fmt"

/*
通过引用或地址传参,在函数调用时可以改变其值

go语言允许向函数传递指针,只需要在函数定义的
参数上设置为指针类型即可.

*/

func main() {
	var a int = 100
	var b int = 1000

	fmt.Printf("交换前 a 的值 : %d\n", a)
	fmt.Printf("交换前 b 的值 : %d\n", b)

	// 调用函数用于交换值
	// &a 指向a变量的地址
	// &b 指向b变量的地址
	swap(&a, &b)
	fmt.Printf("交换后 a 的值: %d\n", a)
	fmt.Printf("交换后 b 的值: %d\n", b)

}

func swap(x *int, y *int) {
	var temp int
	temp = *x // 保存x地址的值
	*x = *y   // 将 y 赋值给 x
	*y = temp // 将temp 赋值给 y
}
