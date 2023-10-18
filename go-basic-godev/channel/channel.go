package main

import (
	"fmt"
	"time"
)

// 只能收数据的channel, 创建一个channel， 顺带创建一个goroutine
func createWorker(id int) chan<- int {
	// 打印乱序是因为，Printf进行了io操作， goroutine会进行调度
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()
	// 上面要开一个 Goroutines-channel， 不然就死循环在这里了，不可能可以return出去的
	return c
}

func worker(id int, c chan int) {
	// 判断 channel是否close还可以用for n := range c
	// 等到 c发完就跳出去
	for n := range c {
		//n, ok := <-c
		//if !ok {
		//	break
		//}
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func chanDemo() {
	//var c chan int // c == nil
	//c := make(chan int)
	// 开一个协程， 让一个人不断地收数据
	//go worker(0, c)
	//c <- 1
	//c <- 2

	// 开10个Worker
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		//channels[i] = make(chan int)
		//go worker(i, channels[i])
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	//c := make(chan int)
	//go func() {
	//	for {
	//		fmt.Printf("%d\n", <-c)
	//	}
	//}()

	/*
			在给 channel发数据之前， 一定要有一个人在收
		    但是在发完1之后， 就必须有人来收， 就比较耗费资源，那么就可以加入一个缓冲区
	*/
	c := make(chan int, 3) // 跟了一个buffer之后， 就有缓冲区了
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond)
}

// channels Close永远都是发送方来Close， 来告诉接收方，我没有新的数据发了
func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c) //告诉接收方， 我发完了
	// close了之后， 接收方还要收 1毫秒时间的空串
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("明明小写字母先发的，但为啥有时候大写字母在前面呢？")
	fmt.Printf("那是因为io操作有一个等待的过程，不会主动交出控制权！！！\n\n")

	fmt.Println("Buffered channel")
	bufferedChannel()

	fmt.Println("Channel close and range")
	channelClose()
}
