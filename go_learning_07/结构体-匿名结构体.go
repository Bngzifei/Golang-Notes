package main

import "fmt"

/*
在定义一些临时数据结构等场景下还可以使用匿名结构体
*/

func main() {
	var user struct {
		Name string
		Age  int
	}
	user.Name = "pprof.cn"
	user.Age = 18
	fmt.Printf("%#v\n", user)
	fmt.Printf("%T\n", user)
}
