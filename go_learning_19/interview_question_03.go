package main

import "fmt"

// [0 0 0 0 0 1 2 3]
//func main() {
//	s := make([]int, 5)
//	s = append(s, 1, 2, 3)
//	fmt.Println(s)
//}

// [1 2 3 4]
func main() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}


/*
new()和make()的区别?

new(T)和make(T,args)是go语言内建函数,用来分配内存,但
适用的类型不同.
new(T)会为T类型的新值分配已置零的内存空间,并返回地址(指针),
即类型为*T的值.换句话说就是,返回一个指针,该指针指向新分配的,
类型为T的零值.适用于值类型,如数组,结构体等.

make(T,args)返回初始化之后的T类型的值,这个值并不是T类型的零值,
也不是指针*T,是经过初始化之后的T的引用.make()只适用于slice map 和 channel



 */