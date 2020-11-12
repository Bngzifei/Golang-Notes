package main

import "fmt"

/*
cap() 函数不适用于map
cap() 函数只适用于array,slice,channel

*/

//func main() {
//	var i interface{}
//	if i == nil {
//		// nil
//		fmt.Println("nil")
//		return
//	}
//	fmt.Println("not nil")
//}

/*
当且仅当接口的动态值和动态类型都为nil时,接口类型值才为nil.
*/

func main() {
	s := make(map[string]int)
	delete(s, "h")
	// 删除map不存在的键值对时,不会报错,相当于没有任何作用,
	// 获取不存在的键值对时,返回值类型对应的零值.
	fmt.Println(s["h"])
}

