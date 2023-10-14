package main

import "fmt"

func slice() {
	s1 := []int{1, 2, 3, 4} //直接初始化了 4 个元素的切片
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))

	s2 := make([]int, 3, 4) //直接初始化了三个元素，容量为 4 的切片
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 7) // 追加一个元素，没有扩容
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 8) // 再追加一个元素，扩容了
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s3 := make([]int, 4) // make 只传入一个参数，表示创建一个 4 个元素的切片
	fmt.Printf("s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))

	// 按照下标索引
	fmt.Printf("s3[2]: %d", s3[2])
	// 超出下标返回，panic
	fmt.Printf("s3[2]: %d", s3[99])
}

func reSlice() {
	s1 := []int{2, 4, 6, 8, 10}
	s2 := s1[1:3]
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s3 := s1[2:] // make 只传入一个参数，表示创建一个 4 个元素的切片
	fmt.Printf("s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))

	s4 := s1[:3] // make 只传入一个参数，表示创建一个 4 个元素的切片
	fmt.Printf("s4: %v, len: %d, cap: %d \n", s4, len(s4), cap(s4))
}

func ShareSlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Printf("share slice s1: %v len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("share slice s2: %v len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2[0] = 99

	fmt.Printf("s2[0]=99 share slice s1: %v len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2[0]=99 share slice s2: %v len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 199)
	fmt.Printf("append s2 share slice s1: %v len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("append s2 share slice s2: %v len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2[0] = 1999
	fmt.Printf("s2[0] = 1999 share slice s1: %v len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2[0] = 1999 share slice s2: %v len: %d, cap: %d \n", s2, len(s2), cap(s2))
}

func sliceExtend() {
	s1 := []int{1, 2, 3, 4} //直接初始化了 4 个元素的切片
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))

	s1 = append(s1, 5)
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
}
