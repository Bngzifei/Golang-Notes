// Go语言中支持匿名函数,即函数可以像普通变量一样被传递或使用,这与C语言的回调函数比较类似.不同的是Go语言支持随时在代码里定义匿名函数

// 匿名函数是指不需要定义函数名的一种函数实现方式,由一个不带函数名的函数声明和函数体组成,下面来具体介绍一下匿名函数的定义及使用.

// 定义一个匿名函数
// 匿名函数的定义格式如下：

// func(参数列表)(返回参数列表){
//     函数体
// }

// 匿名函数的定义就是没有名字的普通函数定义

// 1) 在定义时调用匿名函数
// 匿名函数可以在声明后调用,例如:
func(data int){
    fmt.Println("hello","data")
}(100)
// 注意第三行"}"后的(100),表示对匿名函数进行调用,传递参数为100

// 2)将匿名函数赋值给变量
// 匿名函数体可以被赋值,例如:
// 将匿名函数体保存到f()中
f := func(data int) {
    fmt.Println("hello",data)
}
// 使用f()调用
f(100)

// 匿名函数的用途非常广泛,匿名函数本身是一种值,可以方便的保存在各种容器中实现回调函数和操作封装.


// 匿名函数用作回调函数

// 下面的代码实现对切片的遍历操作,遍历中访问每个元素的操作使用匿名函数来实现.用户传入不同的匿名函数体可以实现对元素不同的遍历操作,代码如下:

package main

import (
    "fmt"
)

// 遍历切片的每个元素,通过给定函数进行元素访问
func visit(list []int,f func(int)) {


    for _,v := range list {
        f(v)
    }
    
}

func main() {
    // 使用匿名函数打印切片内容
    visit([]int{1,2,3,4},func(v int) {
       fmt.Println(v)        
    })
}

// 代码说明如下：
// 第 8 行，使用 visit() 函数将整个遍历过程进行封装，当要获取遍历期间的切片值时，只需要给 visit() 传入一个回调参数即可。
// 第 18 行，准备一个整型切片 []int{1,2,3,4} 传入 visit() 函数作为遍历的数据。
// 第 19～20 行，定义了一个匿名函数，作用是将遍历的每个值打印出来。

// 匿名函数作为回调函数的设计在Go语言的系统包中也比较常见,strings包中就有如下代码:

func TrimFunc(s string, f func(rune) bool) string {
    return TrimRightFunc(TrimLeftFunc(s, f), f)
}

// 使用匿名函数实现操作封装

// 下面这段代码将匿名函数作为map的键值,通过命令行参数动态调用匿名函数,代码如下:

package main

import (
    "fmt"
    "flag"
)


var skillParam = flag.String("skill","","skill to perform")

func main() {
    flag.Parse()

    var skill = map[string]func(){
        "fire":func() {
            fmt.Println("chicken run")
        },
         "run":func() {
            fmt.Println("soldier run")
        },
         "fly":func() {
            fmt.Println("angle run")
        },
    }
    
    if f,ok := skill[*skillParam];ok{
        f()
    }else {
      fmt.Println("skill not found")
    }

}
