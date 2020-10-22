package main

import "fmt"

/*
删除map类型的结构体
 */

type Student struct {
	id int
	name string
	age int
}

func main() {
	// map[键类型]值类型,所以 map[1] = Student{}
	ce := make(map[int]Student)
	ce[1] = Student{
		id:   1,
		name: "老李",
		age:  16,
	}
	ce[2] = Student{
		id:   2,
		name: "老张",
		age:  26,
	}
	// map[1:{1 老李 16} 2:{2 老张 26}]
	fmt.Println(ce)
	// 删除 key = 2 的键值对
	delete(ce,2)
	// map[1:{1 老李 16}]
	fmt.Println(ce)

}