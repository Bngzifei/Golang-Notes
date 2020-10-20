package main

import "fmt"

/*
结构体允许其成员在声明时没有字段名而只有类型,这种没有名字
的字段就称为匿名字段.
*/

type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		"pprof.cn",
		18,
	}

	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.string, p1.int)
}


/*
匿名字段默认采用类型名作为字段名,结构体要求字段名称必须
唯一,因此一个结构体中同种类型的匿名字段只能有一个
 */