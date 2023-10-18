package main

import (
	"fmt"
	"time"
)

type signal struct{}

func worker() {
	fmt.Println("worker is working...")
	time.Sleep(time.Second * 1) // 模拟耗时操作
}

func spawn(f func()) chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to work...")
		f()
		c <- signal{} // 发送信号
	}()
	return c
}

func main() {
	fmt.Println("start a worker...")
	c := spawn(worker)
	<-c // 阻塞等待
	fmt.Println("worker is done!")
}
