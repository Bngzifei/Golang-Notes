package main

import "fmt"

/*
空接口是指没有定义任何方法的接口.因此
任何类型都实现了空接口
空接口类型的变量可以存储任意类型的变量.
*/

func main() {
	// 定义一个空接口
	var x interface{}
	s := "pprof.cn"
	x = s
	fmt.Printf("type:%T value%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}

/*
输出结果:
type:string valuepprof.cn
type:int value:100
type:bool value:true
 */

/*
因为空接口可以存储任意类型值的特点,所以空接口在go语言中的使用
十分广泛.

关于接口需要注意的是,只有当有两个或两个以上的具体类型必须以相同
的方式进行处理时才需要定义接口,不要为了接口而写接口,那样只会增加
不必要的抽象,导致不必要的运行时损耗.
 */