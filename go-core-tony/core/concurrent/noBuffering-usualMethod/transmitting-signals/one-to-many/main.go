package main

import (
	"fmt"
	"sync"
	"time"
)

type signal struct{}

func worker(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d: works done\n", i)
}

func spawnGroup(f func(i int), num int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			<-groupSignal // wait for signal
			fmt.Printf("worker %d: start to work...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		//c <- signal{}
		c <- signal(struct{}{}) // struct{}{} 是 struct{}类型的值
	}()

	return c
}

func main() {
	fmt.Println("start a group of workers...")
	groupSignal := make(chan signal)
	c := spawnGroup(worker, 5, groupSignal)
	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")
	// 关闭一个无缓冲 channel 会让所有阻塞在这个 channel 上的接收操作返回，从而实现了一种 1 对 n 的“广播”机制。
	close(groupSignal) // send signal to start workers
	<-c
	fmt.Println("the group of workers work done!")
}
