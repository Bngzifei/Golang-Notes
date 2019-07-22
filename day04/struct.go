// go语言中数组可以存储同一类型的数据,但在结构体中我们可以为不同项定义不同的数据类型
// 结构体是由一系列具有相同类型 或 不同类型的数据构成的数据集合

// 结构体表示一项纪录,比如保存图书馆的书籍纪录,每本书有以下属性:
// title:标题
// author:作者
// subject:学科
// id:书籍id

// 结构体定义需要使用type和struct语句.struct语句定义一个新的数据类型,结构体中有一个或多个成员.type语句设定了结构体的名称.
// 结构体的格式如下:
// type struct_var struct {
//     member xxx;
// }

package main

import (
	"fmt"
)

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	// 创建一个新的结构体
	fmt.Println(Books{"Go语言", "www.runoob.com", "Go 语言教程", 6495407})
	// 也可以使用key => value 格式
	fmt.Println(Books{title: "Go语言", author: "www.runoob.com", subject: "Go语言教程", book_id: 6495407})
	// 忽略的字段为0 或空
	fmt.Println(Books{title: "Go语言", author: "www.runoob.com"})
}
