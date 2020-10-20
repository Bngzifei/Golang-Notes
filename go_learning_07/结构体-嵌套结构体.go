package main

import "fmt"

/*
一个结构体中可以嵌套包含另一个结构体或结构体指针

*/

//Address地址结构体
type Address struct {
	Province string
	City     string
}

// User 结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

func main() {
	user1 := User{
		Name:   "pprof",
		Gender: "女",
		Address: Address{
			Province: "黑龙江",
			City:     "哈尔滨",
		},
	}

	fmt.Printf("user1=%#v\n", user1)
}


/*
user1=main.User{Name:"pprof", Gender:"女", Address:main.Address{Province:"黑龙江", City:"哈尔滨"}}
 */