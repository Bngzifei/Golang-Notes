package main

import "fmt"

/*

go语言中数组可以存储同一类型的数据,但在结构体中我们
可以为不同项定义不同的数据类型.

结构体是由一系列具有相同类型或不同类型的数据构成的数据
集合.

结构体表示一项记录,比如保存图书馆的书籍记录,每本书有
以下属性:
	title: 标题
	author: 作者
	subject: 科目
	ID:      书籍ID

定义结构体需要使用type和struct语句,struct语句定义一个新的数据类型,
结构体中有一个或多个成员,type语句设定了结构体的名称.
*/

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {

	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	// 忽略的字段为0或空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})

}

/*
{Go 语言 www.runoob.com Go 语言教程 6495407}
{Go 语言 www.runoob.com Go 语言教程 6495407}
{Go 语言 www.runoob.com  0}
*/
