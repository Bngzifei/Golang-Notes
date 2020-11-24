package main

import "fmt"

type Person struct {
	age int
}

func main() {
	person := &Person{28}

	// 1.person.age 此时是将28当做defer函数的参数,会把28缓存在栈中,
	// 等到最后执行该defer语句的时候取出,即输出28
	defer fmt.Println(person.age)
	// 2.defer 缓存的是结构体Person{28}的地址,最终Person{28}的age被重新赋值为29,
	// 所以defer语句最后执行的时候,依靠缓存的地址取出的age便是29,即输出29.
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)
	// 3.闭包引用,输出29
	defer func() {
		fmt.Println(person.age)
	}()
	person.age = 29

	// 又由于defer的执行顺序为先进后出,即 3 2 1 ,所以输出29 29 28

}
