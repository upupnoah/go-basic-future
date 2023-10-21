package main

import "fmt"

func swap1(a, b *int) {
	*b, *a = *a, *b
}

func swap2(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func main() {
	// 指针
	// var a int = 2
	// var pa *int = &a
	// *pa = 3
	// fmt.Println(a)
	a, b, p, q := 1, 2, 3, 4
	swap1(&p, &q)
	fmt.Println(p, q)

	a, b = swap2(a, b)
	fmt.Println(a, b)

	// 交换两个变量的值
	//a, b = b, a
	//fmt.Println(a, b)
}
