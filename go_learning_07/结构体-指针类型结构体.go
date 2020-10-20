package main

import "fmt"

/*
我们还可以通过使用new关键字对结构体进行实例化,得到的是结构体的地址.
*/

type Person struct {
	name, city string
	age        int8
}

func main() {
	var p2 = new(Person)
	// *main.Person 是 Person类
	fmt.Printf("%T\n", p2)
	// p2=&main.Person{name:"", city:"", age:0}
	fmt.Printf("p2=%#v\n", p2)

	/*
		从打印的结果中我们可以看出p2是一个结构体指针
		需要注意的是在go语言中支持对结构体指针直接
		使用.来访问结构体的成员
	*/
	p2.name = "测试"
	p2.age = 18
	p2.city = "北京"
	//p2=&main.Person{name:"测试", city:"北京", age:18}
	fmt.Printf("p2=%#v\n", p2)

	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了
	// 一次new实例化操作
	p3 := &Person{}
	// *main.Person
	fmt.Printf("%T\n", p3)
	// p3=&main.Person{name:"", city:"", age:0}
	fmt.Printf("p3=%#v\n", p3)
	p3.name = "博客"
	p3.age = 30
	p3.city = "成都"
	// p3=&main.Person{name:"博客", city:"成都", age:30}
	// p3.name = "博客" 其实在底层是(*p3).name = "博客",
	// 这是go语言帮我们实现的语法糖
	fmt.Printf("p3=%#v\n", p3)

}
