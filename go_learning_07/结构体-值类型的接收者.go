package main

import "fmt"

/*
当方法作用于值类型的接收者时,go语言会在代码运行时将接收者的值
复制一份.在值类型接收者的方法中可以获取接收者的成员值,但修改
操作只是针对副本,无法修改接收者变量本身
*/

// Person 结构体
type Person struct {
	name string
	age  int8
}

// NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好go语言!\n", p.name)
}

// SetAge2 设置p的年龄
// 使用值接收者
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("测试", 25)
	p1.Dream()
	fmt.Println(p1.age)
	//  (*p1).SetAge2(30)  就是值 对象本身去 调用 SetAge2方法
	p1.SetAge2(30)
	fmt.Println(p1.age)

}

/*
什么时候应该使用指针类型接收者?
1.需要修改接收者中的值
2.接收者是拷贝代价比较大的大对象
3.保证一致性,如果有某个方法使用了指针接收者,
那么其他的方法也应该使用指针接收者

 */


