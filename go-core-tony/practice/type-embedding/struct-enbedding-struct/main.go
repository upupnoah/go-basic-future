package main

import (
	"fmt"
	"reflect"
)

type T1 struct{}

func (T1) T1M1() {
	println("T1's M1")
}
func (*T1) PT1M2() {
	println("PT1's M2")
}

type T2 struct{}

func (T2) T2M1() {
	println("T2's M1")
}
func (*T2) PT2M2() {
	println("PT2's M2")
}

type T struct {
	T1
	*T2
}

func dumpMethodSet(i interface{}) {
	dynType := reflect.TypeOf(i)
	if dynType == nil {
		fmt.Println("there is no dynamic type!")
	}
	n := dynType.NumMethod()
	if n == 0 {
		fmt.Println("there is no method!")
	}
	for j := 0; j < n; j++ {
		fmt.Println("-", dynType.Method(j).Name)
	}
}

func main() {
	t := T{
		T1: T1{},
		T2: &T2{},
	}

	dumpMethodSet(t)
	fmt.Println("-----------------")
	dumpMethodSet(&t)
}
