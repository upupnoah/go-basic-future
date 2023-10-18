package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		go func(i int) { ch <- i + 1 }(i)
	}
	time.Sleep(time.Second) // 模拟耗时操作
}

func consumer(ch <-chan int) {
	for n := range ch {
		fmt.Println(n)
	}
}

func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		producer(ch)
	}()

	go func() {
		defer wg.Done()
		consumer(ch)
	}()

	wg.Wait()
}
