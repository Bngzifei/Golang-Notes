package main

import "fmt"

/*
理解Go接口要记住一点,接口是一组方法的集合.

接口解除了类型依赖,有助于减少用户的可视方法,屏蔽了内部结构和实现
细节.但是接口实现机制会有运行期开销,也不能滥用接口.相对于包,或者
不会频繁变化的内部模块之间,不需要抽象出接口来强行分离.接口最常用
的使用场景,是对包提供访问,或预留扩展空间.也是体现多态很重要的手段,
说到底,接口的意义:就是解耦合,程序和程序之间的关联程度,降低耦合性.

在面向对象编程中,可以这么说:"接口定义了对象的行为",那么具体的实现
行为就取决于对象了.

在go中,接口是一组方法签名.当一个类型为接口中的所有方法提供定义时,它被
称为实现该接口.它与oop非常相似.接口指定类型应该具有的方法,类型决定
如何实现这些方法.

*/

// 定义一个接口
// 接口就是方法的集合,就是一堆方法合起来
type USB interface {
	start()
	end()
}

// 实现鼠标类
type Mouse struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "鼠标,准备就绪,可以开始工作,可以开始点点点...")
}
func (m Mouse) end() {
	fmt.Println(m.name, "结束工作,可以安全退出...")
}

// 实现U盘类
type FlashDisk struct {
	name string
}

func (f FlashDisk) start() {
	fmt.Println(f.name, "准备开始工作,可以进行数据存储了...")
}

func (f FlashDisk) end() {
	fmt.Println(f.name, "可以弹出...")
}

func testInterface(usb USB) {
	usb.start()
	usb.end()
}

func main() {
	/*
		接口:interface
			方法的描述的集合
		实现类:只要实现了该接口中的所有的方法,那么该类就叫做该接口的实现类

		注意点:
			1.当需要接口类型的对象时,那么可以使用任意实现类对象代替
			2.接口对象不能访问实现类的属性
	*/

	m := Mouse{"罗技小红"}
	fmt.Println(m.name)
	f := FlashDisk{"闪迪64G"}

	// 定义接口类型的变量
	var usb USB

	// 创建该接口的任意实现类对象
	usb = m
	// 接口对象,不能访问实现类的属性
	//fmt.Println(usb.name)

	usb.start()
	usb.end()

	fmt.Println("--------------")
	testInterface(m)
	testInterface(f)

}
