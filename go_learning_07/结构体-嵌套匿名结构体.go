package main

import "fmt"

/*
当访问结构体成员时会先在结构体中查找该字段,
找不到再去匿名结构体中查找.
*/

type Address struct {
	Province string
	City     string
}

type User struct {
	Name    string
	Gender  string
	Address // 匿名结构体
}

func main() {
	var user2 User
	user2.Name = "pprof"
	user2.Gender = "女"
	// 通过  匿名结构体.字段名访问
	user2.Address.Province = "黑龙江"
	// 直接访问匿名结构体的字段名
	user2.City = "哈尔滨"
	// main.User{Name:"pprof", Gender:"女", Address:main.Address{Province:"黑龙江", City:"哈尔滨"}}
	fmt.Printf("user2=%#v\n", user2)
}
