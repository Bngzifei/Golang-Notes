package main

import "fmt"

//func main() {
//	var s1 []int  // nil 切片和nil相等,一般用来表示一个不存在的切片
//	var s2 = []int{} // 空切片和nil不相等,表示一个空的集合
//	if s2 == nil {
//		fmt.Println("yes nil")
//	} else {
//		fmt.Println("no nil")
//	}
//}

//func main() {
//	i := 65
//	fmt.Println(string(i))  // utf8编码中, 十进制数字65对应的符号是A
//}

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	c := Work{3}
	var a A = c
	var b B = c
	// 13 23
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}

/*
接口.一种类型实现多个接口.结构体Work分别实现了接口A B,所以接口
变量a,b调用各自的 方法ShowA() 和 ShowB(), 输出 13, 23
 */