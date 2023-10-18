package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

//func (p *field) print() {
//	fmt.Println(p.name)
//}
// 改为 值接收者，即可解决问题
func (p field) print() {
	fmt.Println(p.name)
}

func main() {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go v.print()
	}
	data2 := []field{{"four"}, {"five"}, {"six"}}

	// 这个迭代完成之后， v是"six"的拷贝
	for _, v := range data2 {
		// for中的v是复用的
		go v.print()
		// go (*field).print(&v) // 会打印出3个"six"
		//go field.print(v)
	}

	// 如果子goroutine在main主goroutine执行到这里的之后再执行， 那么打印的就是复用的v
	time.Sleep(3 * time.Second)
}
