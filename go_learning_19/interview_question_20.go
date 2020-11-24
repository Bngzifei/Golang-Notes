package main

import "fmt"

//func f() {
//	// 函数结束,即return之前才会执行
//	defer fmt.Println("D")
//	// 先执行
//	fmt.Println("F")
//}
//
//func main() {
//	f()
//	fmt.Println("M")
//}

/*
被调用函数里的defer语句在返回之前就会被执行
输出结果:
F
D
M
*/

//type Person struct {
//	age int
//}
//
//func main() {
//	person := &Person{28}
//	// 1.
//	defer fmt.Println("1:", person.age)
//	// 2. defer缓存的是结构体Person{28}的地址,这个地址指向的结构体没有被改变,
//	// 最后defer语句后面的函数执行的时候取出仍是28
//	defer func(p *Person) { fmt.Println("2:", p.age) }(person)
//	// 3.闭包引用, person的值已经被改变,指向结构体Person{29}
//	defer func() { fmt.Println("3:", person.age) }()
//
//	person = &Person{29}
//}

func greet() func() {
	fmt.Println("Hello!")
	// 创建一个函数,并使它本身返回另一个函数,
	// 返回的函数将作为真正的延迟函数.
	// 在defer语句调用父函数后在其上添加额外的括号来延迟
	// 执行返回的子函数...
	return func() { fmt.Println("Bye!") } // this will be deferred
}

/*
这为开发者提供了什么能力? 因为在函数内定义的匿名函数可以访问完整的
词法环境,这意味着在函数中定义的内部函数可以引用该函数的变量.
在下一个示例中看到的,参数变量在measure函数第一次执行和其延迟执行的
子函数内都能访问到:

*/

func main() {
	defer greet()()
	fmt.Println("Some code here...")
}
