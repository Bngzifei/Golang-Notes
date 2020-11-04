package main

import "fmt"

/*
下面定义一个结构体类型和该类型的一个方法
*/

// 结构体
type User struct {
	Name  string
	Email string
}

// 方法
func (u *User) Notify() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}

func main() {
	// 值类型调用方法
	u1 := User{"golang", "golang@golang.com"}
	u1.Notify()

	// 指针类型调用方法
	u2 := User{"go", "go@go.com"}
	u3 := &u2
	u3.Notify()

}

/*
输出结果:
golang : golang@golang.com
go : go@go.com

解释:
首先我们定义了一个叫做User的结构体类型,然后定义了一个
该类型的方法叫做Notify,该方法的接受者是一个User类型的
值.要调用Notify方法我们需要一个User类型的值或者指针

在这个例子中,当我们使用指针时,Go调整和解引用指针使得调用
可以被执行.注意:当接受者不是一个指针时,该方法对应接受者的
值的副本(意思就是即使你使用了指针调用函数,但是函数的接受者
是值类型,所以函数内部操作还是对副本的操作,而不是指针操作)

注意:当接受者是指针时,即使用值类型调用那么函数内部也是对指针
的操作.方法不过是一种特殊的函数,只需将其还原,就知道 receiver T
和 *T 的差别
*/