package main

import (
	"fmt"
	"reflect"
)

/*
反射是指在程序运行期对程序本身进行访问和修改的能力

1.1.1 变量的内在机制
变量包含类型信息和值信息 var arr[10]int arr[0] = 10
类型信息:是静态的元信息,是预先定义好的
值信息:是程序运行过程中动态改变的.

反射的使用
reflect包 封装了反射相关的方法
获取类型信息:reflect.TypeOf ,是静态的
获取值信息:reflect.ValueOf,是动态的

1.1.3 空接口与反射
反射可以在运行时动态获取程序的各种详细信息
反射获取interface类型信息
*/

//func reflect_type(a interface{}) {
//	t := reflect.TypeOf(a)
//	fmt.Println("类型是: ", t)
//	// kind() 可以获取具体类型
//	k := t.Kind()
//	fmt.Println(k)
//	switch k {
//	case reflect.Float64:
//		fmt.Printf("a is float64\n")
//	case reflect.String:
//		fmt.Println("string")
//	}
//}
//
//func main() {
//	var x float64 = 3.4
//	reflect_type(x)
//}

//type User struct {
//	Id   int
//	Name string
//	Age  int
//}
//
//// 绑方法
//func (u User) Hello() {
//	fmt.Println("Hello")
//}
//
//func Poni(o interface{}) {
//	t := reflect.TypeOf(o)
//	fmt.Println("类型: ", t)
//	fmt.Println("字符串类型: ", t.Name())
//	// 获取值
//	v := reflect.ValueOf(o)
//	fmt.Println(v)
//	// 可以获取所有属性
//	// 获取结构体字段个数: t.NumField()
//	for i := 0; i < t.NumField(); i++ {
//		// 取每个字段
//		f := t.Field(i)
//		fmt.Printf("%s :%v", f.Name, f.Type)
//		// 获取字段的值信息
//		// Interface(): 获取字段对应的值
//		val := v.Field(i).Interface()
//		fmt.Println("val :", val)
//	}
//
//	fmt.Println("==============方法================")
//	for i := 0; i < t.NumMethod(); i++ {
//		m := t.NumMethod()
//		fmt.Println(m.Name)
//		fmt.Println(m.Type)
//	}
//}
//
//func main() {
//	u := User{1, "zs", 20}
//	Poni(u)
//}

// 调用方法
//type User struct {
//	Id   int
//	Name string
//	Age  int
//}
//
//func (u User) Hello(name string) {
//	fmt.Println("Hello: ", name)
//}
//
//func main() {
//	u := User{1, "5lmh.com", 20}
//	v := reflect.ValueOf(u)
//	println(v)
//	// 获取方法
//	m := v.MethodByName("Hello")
//	// 构建一些参数
//	args := []reflect.Value{reflect.ValueOf("6666")}
//	// 没参数的情况下: var args2 []reflect.Value
//	// 调用方法,需要传入方法的参数
//	m.Call(args)
//
//}

// 获取字段tag

type Student struct {
	Name string `json:"name1" db:"name2"`
}

func main() {
	var s Student
	v := reflect.ValueOf(&s)
	// 类型
	t := v.Type()
	// 获取字段
	f := t.Elem().Field(0)
	fmt.Println(f.Tag.Get("json"))
	fmt.Println(f.Tag.Get("db"))
}








































