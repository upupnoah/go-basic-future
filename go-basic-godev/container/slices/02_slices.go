package main

import "fmt"

// 传入一个切片
func updateSlice(s []int) {
	s[0] = 100
}

// 传入一个切片
func printArray(s []int) {
	fmt.Println(s[:])
}

func main() {
	// 初始化一个数组
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	// 创建一个切片， 对于arr[2:6] 的 view
	s := arr[2:6] // 2, 3, 4, 5
	fmt.Println("arr[2:6] = ", s)
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	s1 := arr[2:]
	s2 := arr[:]
	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)

	// 修改slice的值相当于修改原数组对应切片位置的值
	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	// 将数组作为切片传入函数
	fmt.Println("Passing arrays as slices into.")
	printArray(arr[:])

	// Reslice
	fmt.Println("Reslice")
	s2 = s2[:5]
	fmt.Println(s2)

	// 对切片本身再取一个slice
	s2 = s2[2:]
	fmt.Println(s2)

	// Slice 的扩展
	arr1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	S1 := arr1[2:6]
	S2 := s1[3:5] // 往后面取， 只有超过原数组的范围，才会报错
	fmt.Println("Slice的扩展...")
	fmt.Println(S1, S2)

	// Slice的容量， len， 以及只能向后扩展
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d\n",
		s2, len(s2), cap(s2))
	fmt.Println(s1[3:6])
	
}
