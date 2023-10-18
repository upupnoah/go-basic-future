package main

import "fmt"

func times(x, y int) int {
	return x * y
}

func partialTimes(x int) func(int) int {
	return func(y int) int {
		return times(x, y)
	}
}

func main() {
	timesTwo := partialTimes(2) // 以高频乘数2位固定乘数的乘法函数
	timesThree := partialTimes(3)
	timesFour := partialTimes(4)
	fmt.Println(timesTwo(5))
	fmt.Println(timesThree(6))
	fmt.Println(timesFour(7))
}
