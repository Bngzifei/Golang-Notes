package main

import (
	"image"
	"sync"
)

/*
在编程的很多场景下我们需要确保某些操作在高并发的场景下
只执行一次,例如只加载一次配置文件,只关闭一次通道等.

go语言中的sync包中提供了一个针对只执行一次场景的解决方案-
sync.Once.

sync.Once 只有一个Do方法,其签名如下:
	func (o *Once) Do(f func()) {}

注意:如果要执行的函数f需要传递参数就需要搭配闭包来使用.

加载配置文件示例:
延迟一个开销很大的初始化操作到真正用到它的时候再执行是一个
很好的实践.因为预先初始化一个变量(比如在init函数中完成初始化)
会增加程序的启动耗时,而且有可能实际执行过程中这个变量没有用上,
那么这个初始化操作就不是必须要做的.我们来看一个例子:

多个goroutine并发调用Icon函数时不是并发安全的,现代的编译器和CPU
可能会在保证每个goroutine都满足串行一致的基础上自由地重排访问
内存的顺序.

在这种情况下就会出现即使判断了icons不是nil也不意味着变量初始化完成了.
考虑到这种情况,我们能想到的办法就是添加互斥锁.保证初始化icons的时候
不会被其他的goroutine操作,但是这样做又会引发性能问题.

使用sync.Once改造的示例代码如下:
*/

var icons map[string]image.Image

var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

// Icon 是并发安全的
func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

/*
sync.Once其实内部包含一个互斥锁和一个布尔值,互斥锁保证布尔值和数据的
安全,而布尔值用来记录初始化是否完成.这样设计就能保证初始化操作的时候是
并发安全的并且初始化操作也不会被执行多次.
 */