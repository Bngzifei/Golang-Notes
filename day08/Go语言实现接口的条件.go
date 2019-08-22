// Go语言实现接口的条件

// 如果一个任意类型T的方法集为一个接口类型的方法集的超集,则我们说类型T实现了此接口类型.T可以是一个非接口类型,也可以是一个接口类型.

// 实现关系在Go语言中是隐式的.两个类型之间的实现关系不需要在代码中显式地表示出来.Go语言中没有类似于implements的关键字.Go编译器将自动在需要的时候检查两个类型之间的实现关系.

// 接口定义后,需要实现接口,调用方才能正确编译通过并使用接口.接口的实现需要遵循两条规则才能让接口可用.

// 接口被实现的条件一:接口的方法与实现接口的类型方法格式一致
// 在类型中添加与接口签名一致的方法就可以实现该方法.签名包括方法中的名称,参数列表,返回参数列表.也就是说,只要实现接口类型中的方法的名称,参数列表,返回参数列表中的任意一项与接口要实现的方法不一致,那么接口的这个方法就不会被实现.

// 为了抽象数据写入的过程,定义DataWriter接口来描述数据写入需要实现的方法,接口中的WriterData()方法表示将数据写入,写入方无须关心写入到哪里.实现接口的类型实现WriteData方法时,会具体编写将数据写入到什么结构中.这里使用file结构体实现DataWriter接口的WriteData方法,方法内部只是打印一个日志,表示有数据写入,详细实现参考下面的代码:

// 数据写入器的抽象:

package main

import "fmt"

// 定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error
}

// 定义文件结构,用于实现DataWriter
type file struct {
}

// 实现DataWriter接口的WriteData方法
func (d *file) WriteData(data interface{}) error {
	// 模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

func main() {
	// 实例化file
	f := new(file)

	// 声明一个DataWriter的接口
	var writer DataWriter

	// 将接口赋值f,也就是*file类型
	writer = f

	// 使用DataWriter接口进行数据写入
	writer.WriteData("data")

}

// 代码说明如下：
// 第 8 行，定义 DataWriter 接口。这个接口只有一个方法，即 WriteData()，输入一个 interface{} 类型的 data，返回一个 error 结构表示可能发生的错误。
// 第 17 行，file 的 WriteData() 方法使用指针接收器。输入一个 interface{} 类型的 data，返回 error。
// 第 27 行，实例化 file 赋值给 f，f 的类型为 *file。
// 第 30 行，声明 DataWriter 类型的 writer 接口变量。
// 第 33 行，将 *file 类型的 f 赋值给 DataWriter 接口的 writer，虽然两个变量类型不一致。但是 writer 是一个接口，且 f 已经完全实现了 DataWriter() 的所有方法，因此赋值是成功的。
// 第 36 行，DataWriter 接口类型的 writer 使用 WriteData() 方法写入一个字符串。

// 运行代码，输出如下：
// WriteData: data

// 1.>函数名不一致导致的报错:
// 在以上代码的基础上尝试修改部分代码,造成编译错误,通过编译器的报错理解如何实现接口的方法.首先,修改file结构的WriteData()方法名,蒋政额方法签名修改为...
// func (d *file) WriteDataX(data interface{}) error {
// 编译代码，报错：
// cannot use f (type *file) as type DataWriter in assignment:
//         *file does not implement DataWriter (missing WriteData method)

// 报错的位置在第 33 行。报错含义是：不能将 f 变量（类型*file）视为 DataWriter 进行赋值。原因：*file 类型未实现 DataWriter 接口（丢失 WriteData 方法）。

// WriteDataX方法的签名本身是合法的.但编译器扫描到底33行时,发现尝试将*file类型赋值给DataWritershih ,需要检查类型是否完全实现了DataWriter接口.显然,编译器因为没有找打DataWriter需要的WriteData方法而报错.

// 2.>实现接口的方法签名不一致导致的报错
// 将修改的代码恢复后,再尝试修改Writer

// 到这里了
http://c.biancheng.net/view/78.html