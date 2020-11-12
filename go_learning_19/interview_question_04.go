package main

import "fmt"

//func main() {
//	s1 := []int{1, 2, 3}
//	s2 := []int{4, 5}
//	// 将一个切片追加到另外一个切片上  append(s1,s2…)
//	s1 = append(s1, s2...)
//	fmt.Println(s1)
//}


var (
	// 不能通过编译.
	// 简短模式变量声明:
	// 必须使用显示初始化
	// 不能提供数据类型,编译器会自动推导
	// 只能在函数内部使用简短模式
	size := 1024
	max_size = size * 2
)

func main() {
	fmt.Println(size, max_size)
}
