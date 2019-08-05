go语言的类型或结构体没有构造函数的功能.结构体的初始化过程可以使用函数封装实现.


其他编程语言构造函数的一些常见功能及特性如下:
1.每个类可以添加构造函数,多个构造函数使用函数重载实现
2.构造函数一般与类名同名,且没有返回值
3.构造函数有一个静态构造函数,一般用这个特性来调用父类的构造函数.
4.对于C++来说,还有默认构造函数,拷贝构造函数等.

多种方式创建和初始化结构体----模拟构造函数重载

如果使用结构体描述猫的特性,那么根据猫的颜色和名字可以有不同种类的猫.那么不同的颜色和名字就是结构体的字段,同时可以使用颜色和名字构造不同种类的猫的实例,这个过程可以参考下面的代码:

type Cat struct {
    Color string
    Name string
}

func NewCatByName(name string) *Cat {
    return &Cat{
        Name:name,
    }
}

func NewCatByColor(color string) *Cat {
    return &Cat{
        Color:color
    }
}

// 代码说明如下:
// 第1行定义Cat结构,包含颜色和名字字段
// 第6行定义用名字构造猫结构的函数,返回Cat指针
// 第7行取地址实例化猫的结构体
// 第8行初始化猫的名字字段,忽略颜色字段
// 第12行定义用颜色构造猫结构的函数,返回Cat指针


// 在这个例子中,颜色和名字属于两个属性都是字符串.由于go语言中没有函数重载,为了避免函数名字冲突,使用NewCatByName()和NewCatByColor()两个不同的函数名表示不同的Cat构造过程.

// 带有父子关系的结构体的构造和初始化---模拟父级构造调用

// 黑猫是一种猫,猫是黑猫的一种泛称.同时描述这两种概念时,就是派生,黑猫派生自猫的种类.使用结构体描述猫和黑猫的关系时,将猫(Cat)的结构体嵌入到黑猫(BlackCat)中,表示黑猫拥有猫的特性,然后再使用两个不同的构造函数分别构造出黑猫和猫两个结构体实例,参考下面的代码:


type Cat struct {
    Color string
    Name string
}

type BlackCat struct {
    Cat  //嵌入Cat,类似于派生
}

// "构造基类"
func NewCat (name string) *Cat {
    return &Cat {
        Name:name,
    }
}

// "构造子类"
func NewBlackCat(color string) *BlackCat {
    cat := &BlackCat{}
    cat.Color = color
    return cat
}

// 代码说明如下:
第6行,定义BlackCat结构,并嵌入了Cat结构体.BlackCat拥有Cat的所有成员,实例化后可以自由访问Cat的所有成员

第11行,NewCat()函数定义了Cat的构造过程,使用名字作为参数,填充Cat结构体

第18行,NewBlackCat()使用color作为参数,构造返回BlackCat指针.

第19行,实例化BlackCat结构,此时Cat也同时被实例化
第20行,填充BlackCat中嵌入的Cat颜色属性,BlackCat没有任何成员,所有的成员都来自于Cat.

这个例子中,Cat结构体类似于面向对象中的"基类".BlackCat嵌入Cat结构体,类死于面向对象中的"派生".实例化时,BlackCat中的Cat也会一并被实例化.

总之,go语言中没有提供构造函数相关的特殊机制,用户根据自己的需求,将参数使用函数传递到结构体构造参数中即可完成构造函数的任务.
