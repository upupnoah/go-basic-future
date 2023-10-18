package main

import (
	"fmt"
	"reflect"
)

type T struct{}

type S T // S是一个新的类型，没有T的方法集合
//type S = T // S 和 T 的类型是相同的，包含T的方法集合

func (T) M1() {}
func (T) M2() {}

func (*T) M3() {}
func (*T) M4() {}

// dumpMethodSet 工具函数，打印方法集
func dumpMethodSet(i interface{}) {
	dynType := reflect.TypeOf(i)

	if dynType == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}

	n := dynType.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", dynType)
		return
	}

	fmt.Printf("%s's method set:\n", dynType)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynType.Method(j).Name)
	}
	fmt.Println()
}

func main() {
	var n int
	dumpMethodSet(n)
	dumpMethodSet(&n)

	var t T
	dumpMethodSet(t)
	dumpMethodSet(&t)

	var s S
	dumpMethodSet(s)
	dumpMethodSet(&s)
}