package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	sum := func(a, b int) int {
		return a + b
	}(3, 4)
	fmt.Println(sum)
}
