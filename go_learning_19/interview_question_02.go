package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		// 这是新手常会犯的错误写法,for range 循环
		// 的时候会创建每个元素的副本,而不是元素的引用,
		// 所以 m[key] = &val取的都是变量val的地址,
		// 所以最后map中所有元素的值都是变量val的地址,
		// 因为val最后被赋值为3,所以所有输出都是3
		value := val
		m[key] = &value  // 取的是val的地址
	}

	for k, v := range m {
		fmt.Println(k, "===>", *v)
	}
}


/*
0 -> 3
1 -> 3
2 -> 3
3 -> 3
 */
