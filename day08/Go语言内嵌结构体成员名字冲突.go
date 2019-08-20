// Go语言内嵌结构体成员名字冲突

// 嵌入结构体内部可能拥有相同的成员名,成员名重名时会发生什么?

package main

import (
	"fmt"
)

type A struct {
	a int
}

type B struct {
	a int
}

type C struct {
	A
	B
}

func main() {
	c := &C{}
	// c.A.a = 1
	c.a = 1
	fmt.Println(c)
}

// 编译器告知 C 的选择器 a 引起歧义，也就是说，编译器无法决定将 1 赋给 C 中的 A 还是 B 里的字段 a。

// 在使用内嵌结构体时，Go语言的编译器会非常智能地提醒我们可能发生的歧义和错误。
