package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateChan() chan int {
	ch := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			ch <- i
			i++
		}
	}()
	return ch
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	//var c1, c2 chan int   nil
	c1, c2 := generateChan(), generateChan()
	w := createWorker(0)
	for {
		select {
		case v1 := <-c1:
			//fmt.Println("received from v1:", v1)
			w <- v1
		case v2 := <-c2:
			//fmt.Println("received from v2:", v2)
			w <- v2
			//default:
			//	fmt.Println("no value received")
		}
	}
}
