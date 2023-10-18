package main

import (
	"fmt"
	"reflect"
)

//type MyInt int
//
//func (n *MyInt) Add(m int) {
//	*n = *n + MyInt(m)
//}
//
//type t struct {
//	a, b int
//}
//
//type S struct {
//	*MyInt
//	t
//	io.Reader
//	s string
//	n int
//}

//func main() {
//	m := MyInt(17)
//	r := strings.NewReader("Hello, go!")
//	s := S{
//		MyInt: &m,
//		t: t{
//			a: 1,
//			b: 2,
//		},
//		Reader: r,
//		s:      "demo",
//	}
//	var sl = make([]byte, len("hello, go"))
//	s.Read(sl)
//	fmt.Println(string(sl))
//	s.Add(5)
//	fmt.Println(*(s.MyInt))
//}

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

type I interface {
	M1()
	M2()
}

type T struct {
	I
}

type TT struct {
	II
}

type II I
type III = I

func (T) M3() {}

func main() {
	//var t T
	//var p *T
	//dumpMethodSet(t)
	//dumpMethodSet(p)

	//var i II
	//dumpMethodSet(i)
	//dumpMethodSet(&i)

	var tt TT
	dumpMethodSet(tt)
	dumpMethodSet(&tt)
}
