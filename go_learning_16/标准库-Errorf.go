package main

import "fmt"

/*
Errorf函数根据format参数生成格式化字符串并返回一个包含该字符
串的错误.

func Errorf(format string,a ...interface{}) error

通常使用这种方式来自定义错误类型,例如:
err := fmt.Errorf("这是一个错误")

格式化占位符

*printf系列函数都支持format格式化参数,在这里我们按照占位符
将被替换的变量类型划分,方便查询和记忆.

通用占位符

%v					值的默认格式表示
%+v					类似%v,但输出结构体时会添加字段名
%#v					值的Go语法表示
%T					打印值的类型  类似Python 中的type
%%					百分号

*/

//func main() {
//	fmt.Printf("%v\n", 100)   // 100
//	fmt.Printf("%v\n", false) // false
//	// 定义一个结构体
//	o := struct{ name string }{"枯藤"}
//	fmt.Printf("%v\n", o)  // {枯藤}
//	fmt.Printf("%#v\n", o) // struct { name string }{name:"枯藤"}
//	fmt.Printf("%T\n", o)  // struct { name string }
//	fmt.Printf("100%%\n")  // 100%
//}

/*
宽度标识符
	宽度通过一个紧跟在百分号后面的十进制数指定,如果未指定宽度,则表示值
时除必需之外不作填充.精度通过(可选的)宽度后跟点号后跟的十进制数指定.
如果未指定精度,会使用默认精度;如果点号后没有跟数字,表示精度为0.
*/

//func main() {
//	n := 88.88
//	fmt.Printf("%f\n", n)
//	fmt.Printf("%9f\n", n)
//	fmt.Printf("%.2f\n", n)
//	fmt.Printf("%9.2f\n", n)
//	fmt.Printf("%9.f\n", n)
//}

/*
获取输入:
go语言fmt包下有fmt.Scan  fmt.Scanf  fmt.Scanln三个函数
可以在程序运行过程中从标准输入获取用户的输入

fmt.Scan 函数的签名如下:
	func Scan(a ...interface{}) (n int,err error)

Scan从标准输入扫描文本,读取由空白符分隔的值保存到传递给本函数的参数中,
换行符视为空白符.
本函数返回成功扫描的数据个数和遇到的任何错误.如果读取的数据个数比
提供的参数少,会返回一个错误报告原因.

具体的代码示例如下:
*/

//func main() {
//	var (
//		name    string
//		age     int
//		married bool
//	)
//
//	fmt.Scan(&name, &age, &married)
//	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
//}

/*
fmt.Scanf
函数签名如下:
	func Scanf(foramt string,a ... interface{}) (n int,err error)
Scanf从标准输入扫描文本,根据format参数指定的格式去读取由空白符分隔的值保存到传递给
本函数的参数中.
本函数返回成功扫描的数据个数和遇到的任何错误

*/

func main() {
	var (
		name    string
		age     int
		married bool
	)

	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t  \n", name, age, married)

}
