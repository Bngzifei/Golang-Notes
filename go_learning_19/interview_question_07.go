package main

import "fmt"

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func main() {
	fmt.Println(x, y, z, k, p)
}

/*
B. var x interface{} = nil
D. var x error = nil

nil 只能赋值给指针,chan,func,interface,map或者slice类型的
变量.

 */