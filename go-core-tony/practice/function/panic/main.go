package main

func foo() {
	println("call foo")
	bar()
	println("exit foo")
}

func bar() {
	println("call bar")
	panic("panic occurs in bar")
	zoo()
	println("exit bar")
}

func zoo() {
	println("call zoo")
	println("exit zoo")
}

func main() {
	println("call main")
	foo()
	println("exit main")
}
