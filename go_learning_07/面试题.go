package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*Student)
	students := []Student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}
	for _, stu := range students {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

}
