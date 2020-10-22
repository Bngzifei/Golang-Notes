package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
}

func demo(ce []Student) {
	// 切片是引用传递,是可以改变值的
	ce[1].age = 999
	ce = append(ce, Student{3, "xiaowang", 56})

}

func main() {
	var ce []Student // 定义一个切片类型的结构体
	ce = []Student{
		Student{1, "小王", 22},
		Student{2, "小张", 33},
	}
	// [{1 小王 22} {2 小张 33}]
	fmt.Println(ce)
	// 可以看到是 小张的年龄被修改成了999
	demo(ce)
	// [{1 小王 22} {2 小张 999}]
	fmt.Println(ce)

}
