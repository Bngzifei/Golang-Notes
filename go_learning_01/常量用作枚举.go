package main

import "unsafe"

/*
常量还可以用作枚举:
const (
    Unknown = 0
    Female = 1
    Male = 2
)
数字 0,1,和2分别代表未知性别,女性和男性
常量可以用len(),cap(),unsafe.Sizeof()函数计算表达式的值.
常量表达式中,函数必须是内置函数,否则编译不通过:
*/

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main() {
	println(a, b, c)
}
