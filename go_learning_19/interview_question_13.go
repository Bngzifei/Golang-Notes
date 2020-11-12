package main

import "fmt"

//func hello(i int) {
//	fmt.Println(i)
//}
//
//func main() {
//	i := 5
//	// hello() 函数在执行defer语句的时候会保存一份
//	// 副本,在实际调用hello()函数时用,所以是 5
//	defer hello(i)
//	i = i + 10
//}


type People struct {}


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
	// showA
	// showB
	// 结构体嵌套.
	// Teacher没有自己的ShowA(),所以调用内部类型的People的
	// 同名方法,需要注意的是第5行调用的是People自己的ShowB方法.
	t.ShowA()
}





