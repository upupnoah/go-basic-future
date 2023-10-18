package main

import "fmt"

// sum是闭包中的封闭变了， 有了这个sum，就可以进行累加了
func adder() func(int) int {
	sum := 0 // sum is a closure variable 闭包 自由变量
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	runAdder() // 累加器
}

func runAdder() {
	add := adder()
	fmt.Println("Start runAdder...")
	for i := 0; i < 10; i++ {
		fmt.Println(add(i))
	}
	fmt.Println("runAdder done...")
}
