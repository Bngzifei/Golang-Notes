package main

import "fmt"

func main() {
	var intVal int
	intVal, intVal1 := 1, 2 // 产生编译错误
	fmt.Println(intVal, intVal1)
}

/*
问题记录:
1、 := 符号就是声明(定义)变量的作用,如果 一个变量在:=
声明之前已经使用 var 进行声明过, 那么在 使用 := 声明的时候,如果没有
增加新的变量的 声明 则会报错: no new variables on left side of :=

2、 声明的变量必须被使用，否则报错: intVal1 declared and not used
*/
