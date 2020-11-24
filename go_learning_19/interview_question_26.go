package main

import "fmt"

//const (
//	a = iota
//	b = iota
//)
//
//const (
//	name = "name"
//	c = iota
//	d = iota
//)
//
//
//func main() {
//	fmt.Println(a)
//	fmt.Println(b)
//	fmt.Println(c)
//	fmt.Println(d)
//}

/*
iota是golang语言的常量计数器,只能在常量的表达式中使用.
iota在const关键字出现时将被置为0,const中每新增一行常量
声明将使iota计数一次.
*/

type People interface {
	Show()
}

type Student struct {
}

func (stu *Student) Show() {

}

func main() {
	var s *Student
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	var p People = s
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}
/*
当且仅当动态值和动态类型都为nil时,接口类型值才为nil.
 */