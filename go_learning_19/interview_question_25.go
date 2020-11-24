package main

import "fmt"

//func (i int) PrintInt() {
//	fmt.Println(i)
//}
//
//func main() {
//	var i int = 1
//	i.PrintInt()
//}

/*
基于类型创建的方法必须定义在同一个包内.
上面的代码基于int类型创建了PrintInt()方法,
由于int类型和方法PrintInt()定义在不同的包内,
所以编译出错.
解决的办法可以定义一种新的类型
*/

//type Myint int
//
//func (i Myint) PrintInt() {
//	fmt.Println(i)
//}
//
//func main() {
//	var i Myint = 1
//	i.PrintInt()
//}

type People interface {
	Speak(string) string
}

type Student struct {
}

func (stu *Student) Speak(think string) (talk string) {
	if think == "Speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = &Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
}
