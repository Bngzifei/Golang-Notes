// 在C语言时代大家一般都用过printf()函数,从那个时候开始其实已经在感受可变参数的魅力和价值.如同C语言中的printf()函数,Go语言标准库中的fmt.Println()等函数的实现也严重依赖于语言的可变参数功能.
http://c.biancheng.net/view/60.html

// 本节我们将介绍可变参数的用法,合适地使用可变参数,可以让代码简单易用,尤其是输入输出类函数,比如日志函数等


// 可变参数类型:

// 可变参数是指函数传入的参数个数为不定数量.为了做到这点,首先需要将函数定义为接受可变参数的类型

func myfunc(args...int) {
    for _,arg := range args{
        fmt.Println(arg)
    }
}

// 上面这段代码的意思是,函数myfunc()接受不定数量的参数,这些参数的类型全部是int,所以它可以用以下方式调用.
myfunc(2,3,4)
myfunc(1,2,7,8)

// 形如...type格式的类型只能作为函数的参数类型存在,并且必须是最后一个参数.它是一个语法糖,即这种语法对语言的功能并没有影响,但是更方便程序员使用.通常来说,使用语法糖能够增加程序的可读性,从而减少程序出错的机会


// 从内部实现机理上来说,类型...type本质上是一个数组切片,也就是[]type,这也是为什么上面的参数args可以用for循环来获得每个传入的参数

// 假如没有...type这样的语法糖,开发者将不得不这么写:

func myfunc2(args []int){
    for _,arg := range args{
        fmt.Println(arg)
    }
}

// 从函数的实现角度来看,这没有任何影响,该怎么写就怎么写.但从调用方来说,情形则完全不同:
myfunc2([]int{1,3,7,13})

// 会发现,我们不得不加上[]int{}来构造一个数组切片实例.但是有了...type这个语法糖,我们就不用自己来处理了.


// 可变参数的传递:
// 假设有另外一个变参函数叫做myfunc3(args...int),下面的例子演示了如何向其传递变参:

func myfunc(args...int){
    // 按原样传递
    myfunc3(args...)
    //传递片段,实际上任意的int slice都可以传进去
    myfunc3(args[1:]...)
}

// 任意类型的可变参数
// 之前的例子中将可变参数类型约束为int,如果你希望传任意类型,可以指定类型为interface{}.下面是Go语言标准库中fmt.Printf()的函数原型:

func Printf(format string,args...interface{}){
    //...
}

// 用interface{}传递任意类型数据时Go语言的惯例用法.使用interface{}仍然是类型安全的这和C/C++不一样,下面通过示例来了解一下如何分派传入的interface{}类型的数据.

package main

import "fmt"

func MyPrintf(args...interface{}){
    for _,arg := range args{
        switch arg.(type) {
        case int:
            fmt.Println(arg,"is an int value")
        case string:
            fmt.Println(arg,"is a string value")
        case int64:
            fmt.Println(arg,"is an int64 value")
        default:
            fmt.Println(arg,"is an unknown type")
        }
    }
}



func main() {
    var v1 int = 1
    var v2 int64 = 234
    var v3 string = "hello"
    var v4 float32 = 1.234
    MyPrintf(v1,v2,v3,v4)
}

// 该程序的输出结果为:
// 1 is an int value.
// 234 is an int64 value.
// hello is a string value.
// 1.234 is an unknown type.


// 遍历可变参数列表----获取每一个参数的值
// 可变参数列表的数量不固定,传入的参数是一个切片.如果需要获得每一个参数的具体的值时,可以对可变参数变量进行遍历,参见下面代码:

package main 

import (
    "bytes"
    "fmt"
)

// 定义一个函数,参数数量为0~n,类型约束为字符串
func joinStrings(slist...string) string {
    // 定义一个字节缓冲,快速地连接字符串
    var b bytes.Buffer
    // 遍历可变参数列表slist,类型为[]string
    for _,s := range slist {
        // 将遍历出的字符串连续写入字节数组
        b.WriteString(s)
    }

    // 将连接好的字节数组转换为字符串并输出
    return b.String()
}

func main() {
    // 输入三个字符串,将他们连接成一个字符串
    fmt.Println(joinStrings("pig","and","rat"))
    fmt.Println(joinStrings("hammer","mom","and","hawk"))
}


// 代码输出如下：
// pig and rat
// hammer mom and hawk

// 代码说明如下：
// 第 8 行，定义了一个可变参数的函数，slist 的类型为 []string，每一个参数的类型是 string，也就是说，该函数只接受字符串类型作为参数。
// 第 11 行，bytes.Buffer 在这个例子中的作用类似于 StringBuilder，可以高效地进行字符串连接操作。
// 第 13 行，遍历 slist 可变参数，s 为每个参数的值，类型为 string。
// 第 15 行，将每一个传入参数放到 bytes.Buffer 中。
// 第 19 行，将 bytes.Buffer 中的数据转换为字符串作为函数返回值返回。
// 第 24 行，输入 3 个字符串，使用 joinStrings() 函数将参数连接为字符串输出。
// 第 25 行，输入 4 个字符串，连接后输出。


// 如果要获取可变参数的数量,可以使用len()函数对可变参数变量对应的切片进行长度操作,以获得可变参数的数量

// 获得可变参数类型----获得每一个参数的类型

// 当可变参数为interface{}类型时,可以传入任何类型的值.此时,如果需要获得变量类型,可以通过swich类型分支获得变量的类型.下面的代码演示将一系列不同类型的值传入printTypeValue()函数,该函数将分别为不同的参数打印它们的值和类型的详细描述.


打印类型和值:
package main

import (
    "bytes"
    "fmt"
)


fuc printTypeValue(slist ...interface{}) string {
    // 字节缓冲作为快速字符串连接
    var b bytes.Buffer

    // 遍历参数
    for _,s := range slist {
        // 将interface{}类型格式化为字符串
        str := fmt.Println("%v",s)

        // 类型的字符串描述
        var typeString string

        // 对s进行类型断言
        switch s.(type) {
        case bool://当s为布尔类型时
            typeString = "bool"
        case string://当s为字符串类型时
            typeString = "string"
        case int:  //当s为整型时
            typeString = "int"
        }

        // 写值字符串前缀
        b.WriteString("value: ")

        //写入值
        b.WriteString(str)

        // 写类型前缀
        b.WriteString("type: ")

        // 写类型字符串
        b.WriteString(typeString)
        
        // 写入换行符
        b.WriteString("\n")

    }
    return b.String()
}

func main() {
    // 将不同类型的变量通过printTypeValue()打印出来
    fmt.Println(printTypeValue(100,"str",true))
}


// 代码输出如下：
// value: 100 type: int
// value: str type: string
// value: true type: bool

// 代码说明如下：
// 第 8 行，printTypeValue() 输入不同类型的值并输出类型和值描述。
// 第 11 行，bytes.Buffer 字节缓冲作为快速字符串连接。
// 第 14 行，遍历 slist 的每一个元素，类型为 interface{}。
// 第 17 行，使用 fmt.Sprintf 配合%v动词，可以将 interface{} 格式的任意值转为字符串。
// 第 20 行，声明一个字符串，作为变量的类型名。
// 第 23 行，switch s.(type) 可以对 interface{} 类型进行类型断言，也就是判断变量的实际类型。
// 第 24～29 行为 s 变量可能的类型，将每种类型的对应类型字符串赋值到 typeString 中。
// 第 33～42 行为写输出格式的过程.

// 在多个可变参数函数中传递参数

// 可变参数变量是一个包含所有参数的切片,如果要在多个可变参数中传递参数,可以在传递时在可变参数变量中默认添加...,将切片中的元素进行传递,而不是传递可变参数变量本身.


// 下面的例子模拟print()函数及实际调用的rawPrint()函数,两个函数都拥有可变参数,需要将参数从print传递到rawPrint中.


// 可变参数传递:


package main

import (
    "fmt"
)

// 实际打印的函数
func rawPrint(rawList ... interface {}) {

    // 遍历可变参数切片
    for _,a := range rawList {
        // 打印参数
        fmt.Println(a)
    }
}


// 打印函数封装
func print(slist ...interface {}) {
    // 将slist可变参数切片完整传递给下一个函数
    rawPrint(slist...)
}


func main() {
    print(1,2,3)
}


// 代码输出如下：
// 1
// 2
// 3

// 对代码的说明：
// 第 9～13 行，遍历 rawPrint() 的参数列表 rawList 并打印。
// 第 20 行，将变量在 print 的可变参数列表中添加...后传递给 rawPrint()。
// 第 25 行，传入 1、2、3 这 3 个整型值进行打印。


// 如果尝试将第20行修改为:
// rawPrint("fmt",slist)
// 再次执行代码,将输出:
// [1,2,3]

// 此时,slist(类型为[]interface{})将被作为一个整体传入rawPrint(),rawPrint()函数中遍历的变量也就是slist的切片值

// 可变参数使用...进行传递与切片键使用append连接是同一个特性












