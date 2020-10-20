package main

import "fmt"

type Test struct {
	a,b,c,d int8
}

func main() {
	n := Test {
		1,2,3,4,
	}

	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
}




