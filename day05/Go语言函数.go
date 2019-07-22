// 函数是组织好的,可重复使用的,用来实现单一或相关联的代码段,其可以提高应用的模块性和代码的重复利用率.

// Go语言支持普通函数,匿名函数和闭包,从设计上对函数进行了优化和改进,让函数使用起来更加方便.

// Go语言的函数属于一等公民(first-class),也就是说:
// 1.函数本身可以作为值进行传递
// 2.支持匿名函数和闭包
// 3.函数可以满足接口

// 函数构成代码执行的逻辑结构.在Go语言中,函数的基本组成为:关键字func,函数名,参数列表,返回值,函数体和返回语句

// 每一个程序都包含很多的函数:函数是基本的代码块

// 因为Go语言是编译型语言,所以函数编写的顺序是无关紧要的;鉴于可读性的需求,最好把main()函数写在文件的前面,其他函数按照一定的逻辑顺序进行编写(例如函数被调用的顺序)

// 编写多个函数的主要目的是将一个需要很多行代码的复杂问题分解为一系列简单的任务(那就是函数)来解决.而且,同一个任务(函数)可以被调用很多次,有助于代码重用.
//事实上,好的程序是非常注意DRY原则的即不要重复你自己(Don't Repeat Yourself),意思是执行特定任务的代码只能在程序里面出现一次.

// 当函数执行到代码块的最后一行( }之前 )或者return语句的时候会退出,其中return语句可以带有零个或多个参数;这些参数将作为返回值供调用者使用.简单的return语句也可以用来结束for死循环,或者结束一个协程(goroutine)


// Go语言里面拥有三种类型的函数:
// 1）普通的带有名字的函数
// 2）匿名函数或者lambda函数
// 3) 方法

// 普通函数声明(定义)
// 函数声明包括函数名,形式参数列表,返回值列表(可以省略这部分)以及函数体

// func 函数名(形式参数列表)(返回值列表){
//     函数体
// }

// 形式参数列表描述了函数的参数名以及参数类型.这些参数作为局部变量,其值由参数调用者提供.返回值列表描述了函数返回值的变量名以及类型.如果函数返回一个无名变量或者没有返回值,返回值列表的括号是可以省略的.

// 如果一个函数声明不包括返回值列表,那么函数体执行完毕后,不会返回任何值.例如在下面的hypot函数中:

func hypot(x,y float64) float64 {
    return math.Sqrt(x*x + y*y)
}
fmt.Println(hypot(3,4)) // "5"

// x和y是形参名,3和4是调用时候传入的实际参数,函数返回了一个float64类型的值.返回值也可以像形参一样被命名.在这种情况下,每个返回值被声明成一个局部变量,并根据该返回值的类型,将其初始化为0

// 如果一个函数在声明时,包含返回值列表,该函数必须以return语句结尾.除非函数明显无法运行到结尾处.例如函数在结尾时候调用了panic异常或函数中存在无限循环.

// 正如hypot函数一样,如果一组形参或返回值有相同的类型,我们不必为每个形参都写出参数类型.下面两个声明是等价的:

// 1. func f(i,j,k int,s,t string) {/*...*/}
// 2. func f(i int,j int,k int,s string,t string) {/*...*/}

// 下面,我们给出4种方法声明拥有2个int类型参数和1个int类型返回值的函数,空白标识符_可以强调某个参数未被使用


func add(x int, y int) int {return x + y}
func sub(x, y int) (z int) { z = x - y; return}
func first(x int, _ int) int { return x }
func zero(int, int) int { return 0 }
fmt.Printf("%T\n", add) // "func(int, int) int"
fmt.Printf("%T\n", sub) // "func(int, int) int"
fmt.Printf("%T\n", first) // "func(int, int) int"
fmt.Printf("%T\n", zero) // "func(int, int) int"

// 函数的类型被称为函数的标识符.如果两个函数形式参数列表和返回值列表中的变量类型一一对应,那么这两个函数被认为有相同的类型和标识符.形参和返回值的变量名不影响函数标识符也不影响它们是否可以以省略参数类型的形式表示.

// 每一次函数调用都必须按照声明顺序为所有参数提供实参(参数值).在函数调用时,Go语言没有默认参数值,也没有任何方法可以通过参数名指定形参,因此形参和返回值的变量名对于函数调用者而言没有意义.(这意思就是形参可以随意取名,只是起一个占位的作用,但是一般而言,为了代码可读性,形参都是见名知义的)

// 在函数体中,函数的形参作为局部变量,被初始化为调用者提供的值.函数的形参和有名返回值作为函数最外层的局部变量,被存储在相同的语法块中.实参通过值的方式传递,因此函数的形参是实参的拷贝.对形参进行修改不会影响实参.但是如果实参包括引用类型,如指针,切片(slice)map,function,channel等类型,实参可能会由于函数的间接引用被修改.


// 函数的返回值
// Go语言支持多返回值,多返回值能方便的获得函数执行后的多个返回参数.Go语言经常使用多返回值中的最后一个返回参数返回函数执行中可能发生的错误.示例代码如下:

conn,err := connectToNetwork()

// 在这段代码中,connectToNetwork返回两个参数,conn表示连接的对象,err返回错误

// 其他编程语言中函数的返回值

// 1.) C/C++语言中只支持一个返回值,在需要多个数值时,需要使用结构体返回结果,或者在参数中使用指针变量,然后在函数内部修改外部传入的变量值实现返回计算结果.C++语言中为了安全性,建议在参数返回数据时使用"引用"替代指针.
// 2.) C#语言也没有多返回值特性.C#语言后期加入的ref和out关键字能够通过函数的调用参数获得函数体中修改的数据.
// 3.) lua语言没有指针,但支持多返回值,在大块数据使用时方便很多

// Go语言既支持安全指针,也支持多返回值,因此在使用函数进行逻辑编写时更为方便.


// 1)同一种类型的返回值

// 如果返回值是同一种类型,则用括号将多个返回值类型括起来,用逗号分隔每个返回值的类型.

// 使用return语句返回时,值列表的顺序需要与函数声明的返回值类型一致.

func typedTwoValues() (int,int){
    return 1,2
}
a,b := typedTwoValues()
fmt.Println(a,b)

// 代码输出结果:1 2
// 纯类型的返回值对于代码可读性不是很友好,特别是在同类型的返回值出现时,无法区分每个返回值参数的意义.

// 2)带有变量名的返回值
// Go语言支持对返回值进行命名.这样返回值就和参数一样拥有参数变量名和类型.
// 命名的返回值变量的默认值为类型的默认值,即数值为0,字符串为空字符串,布尔为false,指针为nil等

// 下面代码中的函数拥有两个整型返回值,函数声明时将返回值命名为a何b,因此可以在函数体中直接对函数返回值进行赋值.在命名的返回值方式的函数体中,在函数结束前需要显式地使用return语句进行返回,代码如下:

func namedRetValues() (a,b int){
    a = 1
    b = 2
    return
}

// 提示:
// 同一种类型返回值和命名返回值两种形式只能二选一,混用时将会发生编译错误,例如下面的代码:

func namedRetValues() (a,b int,int)

// 编译报错提示:mixed named and unnamed funxtionparameters
// 意思是:在函数参数中混合使用了命名和非命名参数


// 调用函数
// 函数在定义之后，可以通过调用的方式，让当前代码跳转到被调用的函数中进行执行。调用前的函数局部变量都会被保存起来不会丢失;被调用的函数结束后,恢复到被调用函数的下一行继续执行代码,之前的局部变量也能继续访问.

// 函数内的局部变量只能在函数体中使用,函数调用结束后,这些局部变量都会被释放并且失效.

// G语言的函数调用格式如下:
// 返回值变量列表 = 函数名(参数列表)

// 下面是对各个部分的说明:
// 1)函数名:需要调用的函数名
// 2)参数列表:参数变量以逗号分隔,尾部无须以分号结尾
// 3)返回值变量列表:多个返回值使用逗号分隔

// 例如,加法函数调用样式如下:
result := add(1,1)



























