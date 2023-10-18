package main

import "fmt"

// 数组是值类型
// arr []int 是切片， arr [5]int 才能值类型的
func printArray(arr [10]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println()
}

func printArrayPointer(arr *[10]int) {
	arr[0] = 100 // 不需要 *arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// 遍历方式①
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	// 遍历方式②
	// 使用 range关键字可以同时获得 index 和 value
	// 将 i 替换成 _ 可以只获得值
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// 求和 从 0 1 2 ... 9
	var numbers [10]int
	for i := range numbers {
		numbers[i] = i
	}
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Println(sum)

	/* 数组作为参数传入函数，不能改变原有的数组*/
	fmt.Println("printArray(numbers)")
	printArray(numbers)
	fmt.Println("print the original array")
	for i, v := range numbers {
		fmt.Println(i, v)
	}

	/*  数组指针作为参数， 会修改原有数组的值 */
	fmt.Println("print the array of incoming addresses")
	printArrayPointer(&numbers)
	fmt.Println("print the original array")
	for i, v := range numbers {
		fmt.Println(i, v)
	}

}
