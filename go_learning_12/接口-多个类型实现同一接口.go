package main

import "fmt"

/*
不同的类型还可以实现同一个接口,首先我们定义一个Mover接口,
它要求必须有一个move方法
*/

type Mover interface {
	move()
}

// 例如狗可以动,汽车也可以动,可以使用如下代码实现这个关系:
type dog struct {
	name string
}

type car struct {
	brand string
}

// dog类型实现的Mover接口
func (d dog) move() {
	fmt.Printf("%s会跑\n", d.name)
}

// car类型实现的Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

// 这个时候我们在代码中就可以把狗和汽车当成一个会动的物体来处理了,
// 不再需要关注它们具体是什么,只需要调用它们的move方法就可以了.

func main() {
	var x Mover
	var a = dog{name: "旺财"}
	var b = car{brand: "保时捷"}
	x = a
	x.move()
	x = b
	b.move()
}

/*
并且一个接口的方法,不一定需要由一个类型完全实现,接口的方法
可以通过在类型中嵌入其他类型或者结构体来实现.
 */