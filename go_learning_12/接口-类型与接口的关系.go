package main

import "fmt"

/*
一个类型实现多个接口,而接口间彼此独立,
不知道对方的实现.例如:狗可以叫,也可以动.
我们就分别定义Sayer接口和Mover接口,
如下:

*/
type Sayer interface {
	say()
}

type Mover interface {
	mover()
}

// dog 既可以实现Sayer接口,也可以实现Mover接口
type dog struct {
	name string
}

// 实现Sayer接口
func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d dog) mover() {
	fmt.Printf("%s会动\n", d.name)
}

func main() {
	var x Sayer
	var y Mover

	var a = dog{name: "旺财"}
	x = a
	y = a
	x.say()
	y.mover()
}
