package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, in chan int, done chan bool) {
	for n := range in {
		fmt.Printf("Worker %d received %c\n", id, n)
		// 不要通过共享内存来通信
		// 要用通信来共享内存

		//因为要等多次， 可以并行地向外发，然后外面收
		go func() {
			done <- true
		}() // 通知外面这个事情已经做完了
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	var wg sync.WaitGroup
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
		//<-worker.done // 如果要收一个 workers[i].done的话，就必须先发完，保证了顺序
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-worker.done
	}

	wg.Wait()
	// wait for all of them
	//for _, worker := range workers {
	//	// 因为上面发完了之后又要再发， 那么得先有人收
	//	// 或者在收的时候， 再开一个 Goroutines-channel， 那么就可以继续收了
	//	<-worker.done
	//	<-worker.done
	//}
}

// channels Close永远都是发送方来Close， 来告诉接收方，我没有新的数据发了
func channelClose() {
	worker := createWorker(0)

	go doWorker(0, worker.in, worker.done)
}

func main() {
	chanDemo()
}
