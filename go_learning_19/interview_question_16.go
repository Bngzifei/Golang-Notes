package main

import "fmt"

//func main() {
//	s := [3]int{1, 2, 3}
//	a := s[:0]  // 0 3
//	b := s[:2]  // 2 3
//	c := s[1:2:cap(s)]  // 1 2
//	fmt.Println(a)
//	fmt.Println(b)
//	fmt.Println(c)
//}

/*
数组或切片的截取操作.截取操作有带2个或者3个参数,形如:
[i:j]和[i:j:k],假设截取对象的底层数组长度为L,在操作符
[i:j]中,如果i省略,默认0,如果j省略,默认底层数组的长度,
截取得到的切片长度和容量计算的方法是 j-i,L-i.
操作符[i:j:k], k主要是用来限制切片的容量,但是不能大于
数组的长度L,截取得到的切片长度和容量计算方法是j-i,k-i

*/

//func main() {
//	// 在这里只声明了map m,并没有分配内存空间,不能直接赋值,需要
//	// 使用make(),都提倡使用make() 或者字面量的方式直接初始化map
//	// B处,v,k := m["b"] ,当key为b的元素不存在的时候,v会返回
//	// 值类型对应的零值,k返回false.
//	m := make(map[string]int)
//	m["a"] = 1
//	if v, ok := m["b"]; ok { // B
//		fmt.Println(v)
//	}
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
	fmt.Println(a.ShowB())
	fmt.Println(b.ShowA())
}

/*
接口的静态类型.
a b具有相同的动态类型和动态值,分别是结构体 work和{3};
a 的静态类型是A,b的静态类型是B.接口A不包括方法ShowB(),
接口B也包括方法ShowA(),编译报错.
 */