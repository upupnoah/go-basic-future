package main

import (
	"fmt"
	"github.com/upupnoah/go-basic-future/go-basic/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	//q.Push("QAQ")
	//fmt.Println(q.Pop())
	p := new(int)
	fmt.Printf("%T %v", p, *p)
}
