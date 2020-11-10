package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
)

/*
json是完全独立于语言的文本格式,是k-v的形式
应用场景:前后端交互,系统间数据交互

json使用go语言内置的encoding/json标准库
编码json使用json.Marshal()函数可以对一组数据进行
JSON格式的编码.

func Marshal(v interface{}) ([]byte,error)

*/

//type Person struct {
//	Name  string
//	Hobby string
//}
//
//func main() {
//	p := Person{"5lmh.com", "女"}
//	// 编码json
//	b, err := json.Marshal(p)
//	if err != nil {
//		fmt.Println("json err ", err)
//	}
//	fmt.Println(string(b))
//
//	// 格式化输出
//	b, err = json.MarshalIndent(p, "", "    ")
//	if err != nil {
//		fmt.Println("json err ", err)
//	}
//	fmt.Println(string(b))
//}

// 通过map生成json
//func main() {
//	stu := make(map[string]interface{})
//	stu["name"] = "5lmh.com"
//	stu["age"] = 18
//	stu["sex"] = "man"
//	b, err := json.Marshal(stu)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(b)
//}

/*
解码json使用json.Unmarshal()函数可以对一组数据进行json格式的解码
func Unmarshal(data []byte, v interface{}) error

示例:解析到结构体
*/

//type Person struct {
//	Age       int    `json:"age,string"`
//	Name      string `json:"name"`
//	Niubility bool   `json:"niubility "`
//}
//
//func main() {
//	// 假数据
//	b := []byte(`{"age":"18","name":"5lmh.com","marry":false}`)
//	var p Person
//	err := json.Unmarshal(b, &p)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(p)
//}

/*
解析到interface
*/

//func main() {
//	// 假数据
//	// int 和 float64都当作float64
//	b := []byte(`{"age":1.3,"name":"5lmh.com","marry":false}`)
//
//	// 声明接口
//	var i interface{}
//	err := json.Unmarshal(b, &i)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	// 自动转到map
//	fmt.Println(i)
//	// 可以判断类型
//	m := i.(map[string]interface{})
//	for k, v := range m {
//		switch vv := v.(type) {
//		case float64:
//			fmt.Println(k, "是float64类型", vv)
//		case string:
//			fmt.Println(k, "是string类型", vv)
//		default:
//			fmt.Println("其他")
//		}
//	}
//}

/*
MSGPack
MSGPack是二进制的json,性能更快,更省空间.
需要安装第三方包 go get -u github.com/vmihailenco/msgpack
*/

type Person struct {
	Name string
	Age  int
	Sex  string
}

// 二进制写出
func writerJson(filename string) (err error) {
	var persons []*Person
	// 假数据
	for i := 0; i < 10; i++ {
		p := Person{
			Name: fmt.Sprintf("name%d", i),
			Age:  rand.Intn(100),
			Sex:  "male",
		}
		persons = append(persons, p)
	}
	// 二进制 json序列化
	data, err := msgpack.Marshal(persons)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// 二进制读取
func readJson(filename string) (err error) {
	var persons []*Person
	// 读文件
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 反序列化
	err = msgpack.Unmarshal(data, &persons)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range persons {
		fmt.Printf("%#v\n", v)
	}
	return
}

func main() {
	err := readJson("D:/person.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
}
