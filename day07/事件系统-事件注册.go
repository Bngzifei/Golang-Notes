// 事件系统的基本原理

// 事件系统可以将事件派发者与事件处理者解耦.例如,网络底层可以生产各种事件,在网络连接上后,网络底层只需将事件派发出去,而不需要关心到底哪些代码来响应连接上的逻辑.或者再比如,你注册,关注或者订阅某大V的社交消息后,大V发生的任何事件都会通知你,但他并不用了解粉丝们是如何为她喝彩或者疯狂的.

// 一个事件系统拥有如下特性:
// 1.能够实现事件的一方,可以根据事件ID或者名字注册对应的事件
// 2.事件发起者,会根据注册信息通知这些注册者
// 3.一个事件可以有多个实现方响应

// 通过下面的步骤详细了解事件系统的构成及使用

// 事件注册
// 事件系统需要为外部提供一个注册入口.这个注册入口传入注册的事件名称和对应事件名称的响应函数,事件注册的过程就是将事件名称和响应函数关联并保存起来,详细实现请参考下面代码的RegisterEvent()函数

package main

// 实例化一个通过字符串映射函数切片的map
var eventByName = make(map[string][]func(interface{}))

// 注册事件,提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {
	// 通过名字查找事件列表
	list := eventByName[name]

	// 在列表切片中添加函数
	list = append(list, callback)

	// 将修改的事件列表切片保存回去
	eventByName[name] = list
}

// 调用事件
func CallEvent(name string, param interface{}) {
	// 通过名字找到事件列表
	list := eventByName[name]

	// 遍历这个事件的所有回调
	for _, callback := range list {

		// 传入参数调用回调
		callback(param)

	}

}

//代码说明如下:
// 第4行,创建一个map实例,这个map通过事件名(string)关联回调列表
// ([]func(interface{}),同一个事件名称可能存在多个事件回调,因此使用回调列表保存.回调的函数声明为func(interface{}))

// 第7行,提供给外部的通过事件名注册响应函数的入口
// 第10行,eventByName通过事件名(name)进行查询,返回回调列表([]func(interface{}))
// 第13行,为同一个事件名称在已经注册的事件回调的列表中再添加一个回调函数.
// 第16行,将修改后的函数列表设置到map的对应事件名中.

// 拥有事件名和事件回调函数列表的关联关系后,就需要开始准备事件调用的入口了.

// 事件调用

// 事件调用方和注册方是事件处理中完全不同的两个角色.事件调用方是事发现场,负责将事件和事件发生的参数通过事件系统派发出去,而不关心事件到底由谁处理;事件注册方通过事件系统注册应该响应哪些事件及如何使用回调函数处理这些事件.事件调用的详细实现请参考上面代码的CallEvent()函数

// 代码说明如下:
// 第20行,调用事件的入口,提供事件名称name和参数param.事件的参数表示描述事件具体的细节.例如门打开的事件触发时,参数可以传入谁进来了.
//第23行,通过注册事件回调的eventByName和事件名字查询处理函数列表list

//第26行,遍历这个事件列表,如果没有找到对应的事件,list将是一个空切片

// 第29行,将每个函数回调传入事件参数并调用,就会触发事件实现方的逻辑处理

// 使用事件系统

// 例子中,在main()函数中调用事件系统的CallEvent生成OnSkill事件,这个事件有两个处理函数,一个是角色的OnEvent()方法,还有一个是函数GlobalEvent(),详细代码实现过程请参考下面的代码.


package main

import "fmt"

// 声明角色的结构体
type Actor struct {
}
// 为角色添加一个事件处理函数
func (a*Actor) OnEvent(param interface{}) {
    fmt.Println("actor event:",param)
}


// 全局事件
func GlobalEvent(param interface{}) {
    fmt.Println("global event:",param)
}

func main() {
    // 实例化一个角色
    a := new(Actor)

    //注册名为OnSkill的回调
    RegisterEvent("OnSkill",a.OnEvent)

    // 再次在OnSkill上注册全局事件
    RegisterEvent("OnSkill",GlobalEvent)

    //调用事件,所有注册的同名函数都会被调用
    CallEvent("OnSkill",100)
}





















