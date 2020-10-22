package main

import "fmt"

func foo() (i int) {
	i = 0
	defer func() {
		fmt.Println(i)
	}()
	return 2
}

func main() {
	foo()
}


/*
2
解释:
在有具名返回值的函数中(这里具名返回值为i),
执行return 2的时候实际上已经将 i 的值重新赋值为2.
所以defer closure 输出结果为2而不是1
 */