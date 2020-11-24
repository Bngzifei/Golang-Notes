package main

import "fmt"

//func main() {
//	var b []int  // 声明的是nil切片,这种声明不会分配内存, 优先选择
//	a := []int{}  // 声明的是长度和容量都为0的空切片
//	fmt.Println(a)
//	fmt.Println(b)
//}

//type S struct {
//
//}
//
//func f(x interface{}) {}
//
//func g(x *interface{}) {}
//
//func main() {
//	s := S{}
//	p := &s
//	f(s)
//	g(s)
//	f(p)
//	g(p)
//}

/*
函数参数为interface{}时可以接收任何类型的参数,包括用户自定义类型等,
即使是接收指针类型也用interface{},而不是使用 *interface{}
 */

type S struct {
	m string
}

// 返回的参数是指针类型,所以
func f() *S {
	// 所以可以用&取结构体的指针
	return &S{"foo"}
}


func main() {
	//p := f() 如果*f(),则p是S类型
	//如果填 f(), 则p是*S类型
	p := f()
	fmt.Println(p.m)
}






















