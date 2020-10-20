package main

import "fmt"

// 类型定义
type NewInt int

// 类型别名
type MyInt = int

func main() {
	var a NewInt
	var b MyInt
	// 求 a b 的类型
	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of b:%T\n", b)
}


/*
type of a:main.NewInt
type of b:int
 */

/*
基础数据类型可以表示一些事物的基本属性,但是无法表达事物的全部
或者部分属性时,无法满足

提供了一种自定义数据类型,可以封装多个基本数据类型,这种数据类型
叫  结构体  .
也就是我们可以通过struct 来定义自己的类型了

go 语言中通过struct 来实现面向对象

语言内置的基础数据类型是用来描述一个值的,而结构体是用来描述一组值的.
比如一个人的名字,城市和居住城市等,本质上是一种聚合型的数据类型.


结构体实例化
	只有当结构体实例化时,才会真正地分配内存.也就是必须实例化后才能使用
结构体的字段.

	结构体本身也是一种类型,我们可以像声明内置类型一样使用var关键字声明
结构体类型.
 */
