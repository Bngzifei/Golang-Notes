package main

import "fmt"

//
//import "fmt"
//
//func f(n int) (r int) {
//	defer func() {
//		// r = n + 1
//		r += n
//		recover()
//	}()
//
//	var f func()
//
//	defer f() // 由于此时f()未定义,引发异常,随即执行第一个defer,异常被recover()
//	// 程序正常执行,最后return
//	f = func() {
//		r += 2
//	}
//	return n + 1
//}
//
//func main() {
//	fmt.Println(f(3))
//}

//func main() {
//	var a = [5]int{1, 2, 3, 4, 5}
//	var r [5]int
//	// range 表达式是副本参与循环
//	for i, v := range a {
//		if i == 0 {
//			// {1, 12, 13, 4, 5}
//			a[1] = 12
//			a[2] = 13
//		}
//		// 0-4 1-5
//		r[i] = v
//	}
//	// 原来的a
//	fmt.Println("a = ", a)
//	// 变化后的r
//	fmt.Println("r = ", r)
//}
// range 表达式是副本参与循环,就是说例子中参与循环的是a的副本,
// 而不是真正的a,就这个例子来说,假设b是a的副本,则range循环代码
// 是这样的:

//for i, v := range b {
//	if i == 0{
//		a[1] = 12
//		a[2] = 13
//	}
//	r[i] = v
//}
// 因此无论a被如何修改,其副本b依旧保持原值,并且参与循环的是b,因此v从
// b中取出的仍旧是a的原值,而非修改后的值.

// 如果想要r和a一样输出,修复办法:
func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range &a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

// 修复代码中, 使用*[5]int 作为range表达式,其副本依旧是一个指向原数组a
// 的指针,因此后续所有循环中均是&a指向的原数组亲自参与的,因此v能从
// &a指向的原数组中取出a修改后的值.