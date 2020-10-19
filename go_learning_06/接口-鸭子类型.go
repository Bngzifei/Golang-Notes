package main

import (
	"fmt"
	"math"
)

/*
对于纯面向对象的语言,三大特征之一就是多态,对于go语言,
严格来讲不是基于面向对象的,所以我们可以叫鸭子类型:
一个东西,长得像鸭子,而且会鸭子叫,那它就是鸭子.
*/

// 定义一个接口
type Shape interface {
	// 求周长方法
	peri() float64
	// 求面积方法
	area() float64
}

// 定义实现类:三角形
type Triangle struct {
	a, b, c float64
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	// 半周长
	p := t.peri() / 2
	// 求面积
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}

type Circle struct {
	// 半径
	radius float64
}

func (c Circle) peri() float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

// 测试函数
func testShape(s Shape) {
	fmt.Println("周长：", s.peri(), "面积：", s.area())
}

func main() {
	/*
		多态: 一个事物的多种形态
			go语言:通过接口模拟多态性
			一个实现类的对象:
				看作是一个实现类的类型:能够访问实现类中的方法和属性
				还可以看作是对应的接口类型:只能够访问接口中定义的方法
		接口的用法:
			用法一:一个函数如果接收接口类型作为参数,那么实际上可以传入
		该接口的任意实现类对象作为参数.
			用法二:定义一个类型为接口类型,那么实际上可以赋值任意实现类
		的对象.如果定义了一个接口类型的容器,实际上该容器中可以存储任意的
		实现类对象.
	*/
	var t1 Triangle
	t1 = Triangle{3, 4, 5}
	fmt.Println(t1.peri())
	fmt.Println(t1.area())
	fmt.Println(t1.a, t1.b, t1.c)

	var s1 Shape
	s1 = t1
	fmt.Println(s1.peri())
	fmt.Println(s1.area())

	var c1 Circle
	c1 = Circle{4}
	fmt.Println(c1.peri())
	fmt.Println(c1.area())
	fmt.Println(c1.radius)

	var s2 Shape = Circle{5}
	fmt.Println(s2.peri(), s2.area())

	testShape(t1)

	// 定义一个接口类型的数组
	arr := [4]Shape{t1, s1, c1, s2}
	fmt.Println(arr)

}
