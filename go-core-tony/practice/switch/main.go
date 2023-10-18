package main

type I interface {
	M()
}

type T struct{}

func (T) M() {
}

func main() {
	var x interface{} = 13
	switch x.(type) {
	case nil:
		println("x is nil")
	case int:
		println("the type of x is int")
	case string:
		println("the type of x is string")
	case bool:
		println("the type of x is bool")
	default:
		println("don't support the type")
	}

	var t T
	var i I = t
	switch i.(type) {
	case T:
		println("it is type T")
		// case 只能是实现了接口I的类型
		//case int:
		//	println("it is type int")
		//case string:
		//	println("it is type string")
	}
}
