package main

import "fmt"

/*
可变参函数是指函数的某个参数可有可无,即这个参数个数
可以是0个或者多个
声明可变参数函数的方式是在参数类型前加上 ... 前缀

比如 fmt 包中的Println
func Println(a ...interface{})

函数特征
我们先写一个可变参数函数:

这个函数几乎把可变参数的特征全部表现出来了:
可变参数必须在函数参数列表的尾部,即最后一个(如放前面会引起编译时歧义)
可变参数在函数内部是作为切片来解析的
可变参数可以不填,不填时函数内部当成nil切片处理
可变参数必须是相同类型的(如果需要是不同类型的可以定义为interface{}类型)
*/

func Greeting(prefix string, who ...string) {
	if who == nil {
		fmt.Printf("Nobody to say hi.")
		return
	}
	for _, people := range who {
		fmt.Printf("%s %s\n", prefix, people)
	}
}
// 当没有who参数的时候,输出:
// Nobody to say hi.
// 当有who参数的时候,输出:
// xxx name
// xxx huji
func main() {
	Greeting("xxx","name","huji")
}

/*
调用可变参函数时,可变参部分是可以不传值的,例如:
不传值
传递多个参数
传递切片

总结:
	可变参数必须要位于函数列表尾部
	可变参数是被当做切片来处理的
	函数调用时,可变参数可以不填
	函数调用时,可变参数可以填入切片

 */

