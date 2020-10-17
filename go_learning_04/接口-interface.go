package main

import "fmt"

/*
go语言提供了另外一种数据类型即 接口,
它把所有的具有共性的方法定义在一起,任何
其他类型只要实现了这些方法就是实现了这个接口(这不就是鸭子类型么)
*/

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia,I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone,I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}

/*
I am Nokia,I can call you!
I am iPhone,I can call you!
 */


/*
在上面的例子中,我们定义了一个接口Phone,接口里面有一个
方法call(),然后我们在main函数里面定义了一个Phone类型变量
,并分别为之赋值为NokiaPhone和IPhone,然后调用call方法,
输出结果如下:

I am Nokia,I can call you!
I am iPhone,I can call you!
 */