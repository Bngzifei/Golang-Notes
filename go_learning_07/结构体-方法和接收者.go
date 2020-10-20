package main

import "fmt"

/*
go 语言中的方法(Method)是一种作用于特定类型变量的函数.
这种特定类型的变量叫做 接收者(Receiver),接收者的概念就
类似于其他语言中的this或者self

方法的定义格式如下:
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	函数体
}

其中 :
1. 接收者变量:接收者中的参数变量名在命名时,官方建议使用接收者类型名
的第一个字母小写,而不是self,this之类的命名.
例如，Person类型的接收者变量应该命名为 p，
Connector类型的接收者变量应该命名为c等。

2.接收者类型:接收者类型和参数类似,可以是指针类型和非指针类型

3.方法名,参数列表,返回参数:具体格式与函数定义相同

*/

// Person 结构体
type Person struct {
	name string
	age  int8
}

// NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好go语言!\n", p.name)
}

func main() {
	p1 := NewPerson("测试", 25)
	p1.Dream()
}

/*
方法与函数的区别是:函数不属于任何类型,方法是属于特定的类型
*/
