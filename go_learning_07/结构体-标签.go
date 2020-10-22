package main

import (
	"encoding/json"
	"fmt"
)

/*
Tag 是结构体的元信息,可以在运行的时候通过反射的机制读取出来.

Tag 在结构体字段的后方定义,由一对反引号包裹起来,具体的格式:
 `key1:"value1" key2:"value2"`

结构体标签由一个或多个键值对组成,键与值使用冒号分隔,值用双引号
括起来.键值对之间使用一个空格分隔.
注意事项: 为结构体编写Tag 时,必须严格遵守键值对的规则,结构体
标签的解析代码的容错能力很差,一旦格式写错,编译和运行时不会提示
任何错误,通过反射因为无法正确取值.例如不要在key和value之间添加
空格.
 */

// 例如我们为Student结构体的每个字段定义json序列化时使用的Tag:

// Student 学生
type Student struct {
	// 通过指定tag实现json序列化该字段时的key
	ID int `json:"id"`
	// json序列化时默认使用字段名作为key
	Gender string
	// 私有不能被json包访问
	Name string
}

func main() {
	s1 := Student{
		ID:     1,
		Gender: "女",
		Name:   "pprof",
	}
	data,err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	//json str:{"id":1,"Gender":"女","Name":"pprof"}
	fmt.Printf("json str:%s\n", data)
}