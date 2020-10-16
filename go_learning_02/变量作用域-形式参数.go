package main

import "fmt"

var a int = 20

func main() {
	var a int = 10
	var b int = 20
	var c int = 0

	fmt.Printf("main() 函数中 a = %d\n", a)
	c = sum(a, b)
	fmt.Printf("main() 函数中 c = %d\n", c)

}

func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)

	return a + b
}

/*
初始化局部和全局变量
默认值值为:

数据类型     初始化默认值
  int           0
  float32       0
  pointer       nil
 */
