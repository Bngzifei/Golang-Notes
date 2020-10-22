package main

import (
	"fmt"
	"sync"
	"time"
)

/*
滥用defer 可能会导致性能问题,尤其是在一个"大循环"里

*/

var lock sync.Mutex

func test() {
	lock.Lock()
	lock.Unlock()
}

func testdefer() {
	lock.Lock()
	defer lock.Unlock()
}

func main() {
	func() {
		t1 := time.Now()
		for i := 0; i < 10000; i++ {
			test()
		}
		elapsed := time.Since(t1)
		fmt.Println("test elapsed: ", elapsed)
	}()
	func() {
		t1 := time.Now()
		for i := 0; i < 10000; i++ {
			testdefer()
		}
		elapsed := time.Since(t1)
		fmt.Println("testdefer elapsed: ", elapsed)
	}()

}
/*
test elapsed:  0s
testdefer elapsed:  976.6µs
 */
