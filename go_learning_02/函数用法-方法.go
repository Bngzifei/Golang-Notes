package main

import "fmt"

/*
go 语言中同时有函数和方法.
一个方法就是一个包含了接受者的函数,接受者可以是命名类型
或者结构体类型的一个值或者是一个指针.
所有给定类型的方法属于该类型的方法集.

再次申明:

函数:是公用的,是所有东西都拥有的
方法:是私有的,是自己定义的结构体对应的功能模块
 */

// 定义结构体
type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}


func (c Circle) getArea() float64 {
	// c.radius 即为Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}
