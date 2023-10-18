package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
}

func main() {
	// 向Slice添加元素
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr =", arr)
	s1 := arr[2:6]       // 2 3 4 5
	s2 := s1[3:5]        // 5 6
	s3 := append(s2, 10) // 5 6 10
	s4 := append(s3, 11) // 5 6 10 11
	s5 := append(s4, 12) // 5 6 10 11 12
	fmt.Println("s3 =", s3)
	fmt.Println("s4 =", s4)
	fmt.Println("s5 =", s5)
	// s4 and s5 no longer view arr.
	fmt.Println("arr =", arr) // 0 1 2 3 4 5 6 10

	// 定义 s 类型为slice，值还没有
	var s []int // Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	// 多种定义切片的方式
	fmt.Println("Creating slice...")
	var snew []int
	ss := []int{2, 4, 6, 8, 10, 12, 16, 20}
	sss := make([]int, 16)
	ssss := make([]int, 10, 32)
	fmt.Println(snew, ss, sss, ssss)
	printSlice(snew)
	printSlice(ss)
	printSlice(sss)
	printSlice(ssss)

	fmt.Println("Copying slice...")
	copy(sss, ss) // 将 ss 拷贝到 sss中
	fmt.Println(sss)

	fmt.Println("Deleting elements from slice")
	fmt.Println("Before:", ss)
	ss = append(ss[:3], ss[4:]...) // 跳过第三个就行
	fmt.Println("After:", ss)

	fmt.Println("Popping from front")
	front := ss[0]
	ss = ss[1:]
	fmt.Println("front =", front, "ss =", ss)

	

	fmt.Println("Popping from back")
	tail := ss[len(ss)-1]
	ss = ss[:len(ss)-1]
	fmt.Println("tail = ", tail, "ss =", ss)
}
