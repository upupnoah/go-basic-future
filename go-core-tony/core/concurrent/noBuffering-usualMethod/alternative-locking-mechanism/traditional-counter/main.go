package main

import (
	"fmt"
	"sync"
)

type counter struct {
	sync.Mutex
	i int
}

var cter counter

func Increase() int {
	cter.Lock()
	defer cter.Unlock()
	cter.i++
	return cter.i
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			v := Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
			wg.Done() // 完成一个goroutine
		}(i) // 执行一个goroutine并将i传递给goroutine
	}
	wg.Wait() // 等待所有goroutine执行完毕
}
