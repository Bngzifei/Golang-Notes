package main

import "fmt"

//func main() {
//	i := -5
//	j := +5
//	// -5 +5
//	// %d 表示输出十进制数字,+表示输出数值的符号.这里不表示取反
//	fmt.Printf("%+d %+d", i, j)
//}

type People struct {
}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	// teacher showB
	// 结构体嵌套,People称为内部类型,Teacher称为外部类型
	// 通过嵌套,内部类型的属性,方法,可以为外部类型所有.
	// 就好像外部类型自己的一样.
	// 外部类型还可以定义自己的属性和方法,甚至可以定义与
	// 内部相同的方法,这样内部类型的方法就会被"屏蔽".
	// 这就是面向对象的继承与重载
	t.ShowB()
}
