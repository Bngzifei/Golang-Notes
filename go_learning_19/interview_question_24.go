package main

import "fmt"

//func main() {
//	m := map[int]string{0:"zero",1:"one"}
//	for k,v := range m {
//		// 遍历map,输出是无序的
//		fmt.Println(k,v)
//	}
//}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2    // 1  ret=3   1 0 3 3
	defer calc("1", a, calc("10", a, b))
	a = 0     // 1  ret=2    2 0 2 2
	defer calc("2", a, calc("20", a, b))
	b = 1
}
/*
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
 */