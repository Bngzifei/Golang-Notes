package main

import "fmt"

/*
go 语言中使用结构体也可以实现其他编程语言中
面向对象的继承.
*/

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s 会动!\n", a.name)
}

// Dog 狗
type Dog struct {
	Feet    int8
	*Animal // 通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet:   4,
		// 注意嵌套的是结构体指针
		Animal: &Animal{name: "大黄"},
	}
	// 大黄会汪汪汪~
	d1.wang()
	// 大黄 会动!
	d1.move()

}

/*
结构体中字段大写开头表示可以公开访问,小写表示私有(仅在定义当前
结构体的包中可访问)
 */

