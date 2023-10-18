package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type signal struct{}

var ready bool

func worker(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(time.Second * 1)
	fmt.Printf("worker %d: works done\n", i)
}

func spawnGroup(f func(i int), num int, mu *sync.Mutex) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			for {
				mu.Lock()
				if !ready { // 如果没有准备好，则等待
					mu.Unlock()
					time.Sleep(100 * time.Millisecond)
					continue
				}
				mu.Unlock()
				fmt.Printf("worker %d: start to work...\n", i)
				f(i)
				wg.Done()
				return
			}
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()
	return c
}

func TestCommonPolling(t *testing.T) {
	fmt.Println("start a group of workers...")
	mu := &sync.Mutex{}
	c := spawnGroup(worker, 5, mu)

	time.Sleep(5 * time.Second) //模拟ready钱的准备工作
	fmt.Println("the group of workers start to work...")

	mu.Lock()
	ready = true
	mu.Unlock()

	<-c
	fmt.Println("hte group of workers work done!")
}

func spawnGroupCond(f func(i int), num int, groupSignal *sync.Cond) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			groupSignal.L.Lock()
			for !ready { // 如果没有准备好，则等待
				groupSignal.Wait() // Wait方法会在解锁前阻塞，直到有其他goroutine调用Signal或者Broadcast
			}
			groupSignal.L.Unlock()
			fmt.Printf("worker %d: start to work...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()
	return c
}

func TestCond(t *testing.T) {
	fmt.Println("start a group of workers...")
	groupSignal := sync.NewCond(&sync.Mutex{})
	c := spawnGroupCond(worker, 5, groupSignal)

	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work....")

	groupSignal.L.Lock()
	ready = true
	groupSignal.Broadcast()
	groupSignal.L.Unlock()

	<-c
	fmt.Println("hte group of workers work done!")
}
