package main

import "fmt"

func foo1(x *int) func() {
	return func() {
		*x = *x + 10
		fmt.Printf("foo1 val = %d\n", *x)
	}
}

func foo2(x int) func() {
	return func() {
		x = x + 1
		fmt.Printf("foo2 val = %d\n", x)
	}
}

func foo3() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fmt.Printf("foo3 val = %d\n", val)
	}
}

func show(v interface{}) {
	fmt.Printf("foo4 val = %v\n", v)
}
func foo4() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go show(val)
	}
}

func foo5() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go func() {
			fmt.Printf("foo5 val = %v\n", val)
		}()
	}
}

var foo6Chan = make(chan int, 10)

func foo6() {
	for val := range foo6Chan {
		go func() {
			fmt.Printf("foo6 val = %d\n", val)
		}()
	}
}

func foo7(x int) []func() {
	var fs []func()
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fs = append(fs, func() {
			fmt.Printf("foo7 val = %d\n", x+val)
		})
	}
	return fs
}

func main() {
	foo1(new(int))
	foo2(1)
	foo3()
	foo4()
	foo5()
	foo6()
	foo7(1)
}
