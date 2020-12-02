package main

import "fmt"

//type Foo struct {
//	bar string
//}
//
//func main() {
//	s1 := []Foo{
//		{"A"},
//		{"B"},
//		{"C"},
//	}
//	// 长度和容量:是不同的概念, 长度是指有现在已经多少个元素,容量是指可以放多少个
//	s2 := make([]*Foo, len(s1))
//	for i := range s1 {
//		s2[i] = &s1[i]
//	}
//	fmt.Println(s1[0], s1[1], s1[2])
//	fmt.Println(s2[0], s2[1], s2[2])
//}

func main() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		fmt.Println(counter)
		counter++
		fmt.Println(k, v)
		fmt.Println("-----")
	}
	// 循环结束,counter已经变成3了
	fmt.Println("counter is ", counter)
}
// for range map 是无序的,如果第一次循环到A,则输出3,否则输出2