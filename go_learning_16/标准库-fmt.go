package main

import "fmt"

/*
fmt 包实现了类似c语言printf和scanf的格式化I/O.主要分为向
外输出内容和获取输入内容两大部分.

1.1 向外输出
标准库fmt提供了以下几种输出相关函数

Print 系列函数会将内容输出到系统的标准输出,区别在于Print函数直接
输出内容,Printf函数支持格式化输出字符串,Println函数会在输出内容的
结尾添加一个换行符.

*/

//func main() {
//	// 直接输出内容
//	fmt.Print("在终端打印该信息.")
//	name := "枯藤"
//	// 格式化输出字符串
//	fmt.Printf("我是:%s\n", name)
//	// Println 最后加一个换行符
//	fmt.Println("在终端打印单独一行显示")
//}

/*
Fprint系列函数会将内容输出到一个id.Writer接口类型的变量中w中,我们通常
用这个函数往文件中写入内容.
*/

//func main() {
//	// 向标准输出写入内容
//	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
//	fileObj, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
//	if err != nil {
//		fmt.Println("打开文件出错,err:", err)
//		return
//	}
//	name := "枯藤"
//	// 向打开的文件句柄中写入内容
//	fmt.Fprintf(fileObj, "写入的内容是:%s", name)
//}

/*
注意:只要满足io.Writer接口的类型都支持写入

Sprint系列函数会把传入的数据生成并返回一个字符串.
func Sprint(a ...interface{}) string
func Sprintf(format string, a ... interface{}) string
func Sprintln(a ... interface{}) string

*/

// 简单的示例代码如下:
func main() {
	s1 := fmt.Sprint("枯藤")
	name := "枯藤"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("枯藤")
	fmt.Println(s1, s2, s3)
}
