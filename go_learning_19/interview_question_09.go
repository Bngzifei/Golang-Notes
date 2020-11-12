package main

import (
	"fmt"
	"reflect"
)

//type person struct {
//	name string
//}
//
//func main() {
//	var m map[person]int
//	p := person{"mike"}
//	// 0
//	// 打印一个map中不存在的值时,返回元素类型的零值
//	fmt.Println(m[p])
//}

//func hello(num ...int) {
//	num[0] = 18
//}
//
//func main() {
//	i := []int{5, 6, 7}
//	hello(i...)
//	// 18 就是把索引位置0的值换成了18
//	fmt.Println(i[0])
//}


func main() {
	a := [5]int{1,2,3,4,5}
	t := a[3:4:4]
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(t))
	fmt.Println(t)
}