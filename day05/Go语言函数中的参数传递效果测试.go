// Go语言中传入和返回参数在调用和返回时都使用值传递.这里需要注意的的是指针,切片和map等引用型对象指向的内容在参数传递中不会发生复制,而是将指针进行复制,类似于创建一次引用

package main

import (
	"fmt"
)

// 用于测试值传递效果的结构体
type Data struct {
	complex  []int      // 测试切片在参数传递中的效果
	instance InnerData  //实例分配的innerData
	ptr      *InnerData //将ptr声明为InnerData的指针类型

}

// 代表各种结构体字段
type InnerData struct {
	a int
}

// 值传递测试函数
func passByValue(inFunc Data) Data {
	// 输出参数的成员情况
	fmt.Println("inFunc value:%+v\n", inFunc)

	// 打印inFunc的指针
	fmt.Printf("inFunc ptr:%p\n", &inFunc)

	return inFunc
}

func main() {
	// 准备传入函数的结构
	in := Data{
		complax: []int{1, 2, 3},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{1},
	}
	// 输入结构的成员情况
	fmt.Println("in value:%+v\n", in)

	// 输入结构的指针地址
	fmt.Println("in ptr:%p\n", &in)

	// 传入结构体,返回同类型的结构体
	out := passByValue(in)

	// 输出结构的成员情况
	fmt.Println("out value:%+v\n", out)

	// 输出结构的指针地址
	fmt.Println("out ptr:%p\n", &out)
}

// 1)测试数据类型:
// 为了测试结构体,切片,指针及结构体中嵌套的结构体在值传递中会发生的情况,需要定义一些结构,代码如下:

// 用于测试值传递效果的结构体
type Data struct {
	complax []int //测试切片在参数传递中的效果

	instance InnerData //实例分配的innerData

	ptr *InnerData
}

// 代表各种结构体字段
type InnerData struct {
	a int
}

// 2)值传递的测试函数
// 本节中定义的passByValue()函数用于值传递的测试,该函数的参数和返回值都是Data类型.在调用中,Data的内存会被复制后传入函数,当函数返回时,又会将返回值复制一次,赋给函数返回值的接收变量.代码如下:

// 值传递的测试函数
func passByValue(inFunc Data) {

	// 输出参数的成员情况
	fmt.Printf("inFunc value:%+v\n", inFunc)

	// 打印inFunc的指针
	fmt.Printf("inFunc ptr:%p\n", &inFunc)

	return inFunc
}

// 代码说明如下：
// 第 5 行，使用格式化的%+v动词输出 in 变量的详细结构，以便观察 Data 结构在传递前后的内部数值的变化情况。
// 第 8 行，打印传入参数 inFunc 的指针地址。在计算机中，拥有相同地址且类型相同的变量，表示的是同一块内存区域。
// 第 10 行，将传入的变量作为返回值返回，返回的过程将发生值复制。

// 3)测试流程
// 测试流程会准备一个Data格式的数据结构并填充所有成员,这些成员类型包括切片,结构体成员及指针.通过调用测试函数,传入Data结构数据,并获得返回值,对比输入和输出后的Data结构数值变化,特别是指针变化情况以及输入和输出整块数据是否被复制,代码如下:

// 准备传入函数的结构
in := Data{
    complax: []int{1,2,3},
    instance: InnerData{
        5,
    },

    ptr:&InnerData{1},
}

// 输入结构的成员情况
fmt.Printf("in value: %+v\n",in)

// 输入结构的指针地址
fmt.Printf("in ptr:%p\n",&in)

// 传入结构体,返回同类型的结构体
out := passByValue(in)

// 输出结构的成员情况
fmt.Println("out value:%+v\n",out)

// 输出结构的指针地址
fmt.Printf("out ptr:%p\n",&out)

// 代码说明如下：
// 第 2 行，创建一个 Data 结构的实例 in。
// 第 3 行，将切片数据赋值到 in 的 complax 成员。
// 第 4 行，为 in 的 instance 成员赋值 InnerData 结构的数据。
// 第 8 行，为 in 的 ptr 成员赋值 InnerData 的指针类型数据。
// 第 12 行，打印输入结构的成员情况。
// 第 15 行，打印输入结构的指针地址。
// 第 18 行，传入 in 结构，调用 passByvalue() 测试函数获得 out 返回，此时，passByValue() 函数会打印 in 传入后的数据成员情况。
// 第 21 行，打印返回值 out 变量的成员情况。
// 第 24 行，打印输出结构的地址。

// 运行代码，输出结果为：
// in value: {complax:[1 2 3] instance:{a:5} ptr:0xc042008100}
// in ptr: 0xc042066060
// inFunc value: {complax:[1 2 3] instance:{a:5} ptr:0xc042008100}
// inFunc ptr: 0xc0420660f0
// out value: {complax:[1 2 3] instance:{a:5} ptr:0xc042008100}
// out ptr: 0xc0420660c0

// 从运行结果中发现:
// 所有的Data结构的指针地址发生了变化,意味着所有的结构都是一块新的内存,无论是将Data结构传入函数内部,还是通过函数返回值传回Data都会发生复制行为.

// 所有的Data结构中的成员值都没有发生变化,原样传递,意味着所有参数都是值传递

// Data结构的ptr成员在传递过程中保持一致,表示指针在函数参数值传递中传递的只是指针值,不会复制指针指向的部分.