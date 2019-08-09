// Go语言为任意类型添加方法
// Go语言可以为任何类型添加方法.给一种类型添加方法就像给结构体添加方法一样,因为结构体也是一种类型.

// 1.)为基本类型添加方法
// 在G语言中,使用type关键字可以定义出新的自定义类型.之后就可以为自定义类型添加各种方法.我们习惯于面向过程的方式判断一个值是否为0,例如:
// if v == 0 {
//     // v等于0
// }

// 如果将v当做整型对象,那么判断v值就可增加一个IsZero()方法,通过这个方法就可以判断v值是否为0,例如:
// if v.IsZero() {
//     // v等于0
// }

// 为基本类型添加方法的详细实现流程如下:

package main

import (
	"fmt"
)

// 将inty定义为MyInt类型
type MyInt int

// 为MyInt添加IsZero()方法
func (m MyInt) IsZero() bool {
	return m == 0
}

// 为MyInt添加Add() 方法
func (m MyInt) Add(other int) int {
	return other + int(m)
}

func main() {
	var b MyInt

	fmt.Println(b)

	fmt.Println(b.IsZero())

	b = 1

	fmt.Println(b.Add(2))

}

// 代码说明如下：
// 第 8 行，使用 type MyInt int 将 int 定义为自定义的 MyInt 类型。
// 第 11 行，为 MyInt 类型添加 IsZero() 方法。该方法使用了 (m MyInt) 的非指针接收器。数值类型没有必要使用指针接收器。
// 第 16 行，为 MyInt 类型添加 Add() 方法。
// 第 17 行，由于 m 的类型是 MyInt 类型，但其本身是 int 类型，因此可以将 m 从 MyInt 类型转换为 int 类型再进行计算。
// 第 24 行，调用 b 的 IsZero() 方法。由于使用非指针接收器，b的值会被复制进入 IsZero() 方法进行判断。
// 第 28 行，调用 b 的 Add() 方法。同样也是非指针接收器，结果直接通过 Add() 方法返回。
