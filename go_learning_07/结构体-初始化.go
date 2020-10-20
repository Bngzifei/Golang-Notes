package main

import "fmt"

type Person struct {
	name, city string
	age        int8
}

func main() {
	var p4 Person
	// p4=main.Person{name:"", city:"", age:0}
	fmt.Printf("p4=%#v\n", p4)

	// 使用键值对结构体进行初始化时,键对应结构体的字段,值对应
	// 该字段的初始值
	p5 := Person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}
	//p5=main.Person{name:"pprof.cn", city:"北京", age:18}
	fmt.Printf("p5=%#v\n", p5)

	// 也可以对结构体指针进行键值对初始化
	p6 := &Person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}
	//p6=&main.Person{name:"pprof.cn", city:"北京", age:18}
	fmt.Printf("p6=%#v\n", p6)

	// 当某些字段没有初始值的时候,该字段可以不写.
	// 此时,没有指定初始值的字段的值就是该字段类型的零值
	p7 := &Person{
		city: "北京",
	}
	// p7=&main.Person{name:"", city:"北京", age:0}
	fmt.Printf("p7=%#v\n", p7)

	// 初始化结构体的时候可以简写,也就是初始化的时候不写键,直接写值
	// 使用这种格式初始化时,需要注意:
	// 必须初始化结构体的所有字段
	// 初始值的填充顺序必须与字段在结构体中的声明顺序一致
	// 该方式不能和键值初始化方式混用
	p8 := &Person{
		"pprof.cn",
		"北京",
		18,
	}
	// p8=&main.Person{name:"pprof.cn", city:"北京", age:18}
	fmt.Printf("p8=%#v\n", p8)
}
