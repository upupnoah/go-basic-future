package types

import "io"

type Outer struct {
	Inner
}

type Outer1 struct {
	*Inner
}

type Inner struct {
}

func (o Outer) Name() string {
	return "Outer"
}

func (i Inner) SayHello() {
	println("hello," + i.Name())
}

func (i Inner) Name() string {
	return "Inner"
}

func UseOuter() {
	var o Outer
	o.SayHello()
}

type Outer2 struct {
	// 组合了接口
	io.Closer
}
