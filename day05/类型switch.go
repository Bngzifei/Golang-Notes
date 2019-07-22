package main

import (
	"fmt"
)

// switch 还可以用于获得一个接口变量的动态类型.
// 这种类型swich使用类型断言的语法,在括号中使用关键字type.如果switch在表达式中声明了一个变量,则变量会在每个子句中具有对应的类型.比较符合控制结构语言习惯的方式是在这些case里重用一个名字,实际上是在每个case里声明一个新的变量,其具有相同的名字,但是不同的类型

func functionOfSomeType() {
	fmt.Println("xx")
}

func main() {
	var t interface{}
	t = functionOfSomeType()
	switch t := t.(type) {
	default:
		fmt.Println("unexpected type %T", t) // %T打印任何类型的t
	case bool:
		fmt.Println("boolean %t\n", t) //t的类型是bool
	case int:
		fmt.Println("integer %d\n", t) //t的类型是int
	case *bool:
		fmt.Println("pointer to boolean %t\n", *t) //t的类型是*bool
	case *int:
		fmt.Println("pointer to integer %d\n", *t) //t的类型是*int
	}

}
