package main

import "fmt"

/*
go 语言的结构体没有构造函数,我们可以自己实现.例如;下方的代码
就实现了一个person的构造函数.因为struct是值类型,如果结构体
比较复杂的话,值拷贝性能开销会比较大,所以,该构造函数返回的是
结构体指针类型.
*/

type Person struct {
	name, city string
	age        int8
}

func newPerson(name, city string, age int8) *Person {
	return &Person{
		name: name,
		city: city,
		age:  age,
	}
}

func main() {
	// 调用构造函数
	p9 := newPerson("pprof.cn", "测试", 18)
	// &main.Person{name:"pprof.cn", city:"测试", age:18}
	fmt.Printf("%#v\n", p9)

}
