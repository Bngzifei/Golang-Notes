package main

import "fmt"

/*
常量:
	在程序运行时,不会被修改的量
	数据类型只能是 布尔型,数值型,字符串型
	定义格式  const 常量名 类型=值

可以省略类型说明,因为编译器可以根据变量的值来推断其类型
eg:
	显式类型定义： const b string = "abc"
	隐式类型定义： const b = "abc"

多个相同类型的声明可以简写为:
const c_name1, c_name2 = value1, value2

常量还可以用作枚举:

*/
func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" // 多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)

}
