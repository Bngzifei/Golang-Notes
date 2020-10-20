package main

import "fmt"

/*
指针类型的接收者由一个结构体的指针组成,由于指针的特性,
调用方法时修改接收者指针的任意成员变量,在方法结束后,修改
都是有效的.这种方式就十分接近于其他语言中面向对象时候的
this或者self.例如我们为Person添加一个SetAge方法,来修改
实例变量的年龄.
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

func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// 调用该方法
func main() {
	p1 := NewPerson("测试", 25)
	// 25
	fmt.Println(p1.age)

	p1.SetAge(30)
	// 30
	fmt.Println(p1.age)
}
