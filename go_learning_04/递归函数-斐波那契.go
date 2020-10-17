package main

import "fmt"

/*
Go 语言的递归函数实现斐波那契数列
*/

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}

/*
0	1	1	2	3	5	8	13	21	34
 */