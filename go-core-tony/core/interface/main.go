package main

import "fmt"

type A interface {
	F1()
	F2()
}

type B interface {
	F1()
}

func main() {
	var a A
	var b B
	b = a // A是B的超集或者和B等价的话，就可以这么赋值
	fmt.Println(b)

	fmt.Println("-----------------")

	var aa int64 = 13
	var i interface{} = aa
	v1, ok := i.(int64)                                            // 这种方式， 如果断言失败，那么v1为nil
	fmt.Printf("v1=%d, the type of v1 is %T, ok=%t\n", v1, v1, ok) // v1=13, the type of v1 is int64, ok=true
	v2, ok := i.(string)
	fmt.Printf("v2=%s, the type of v2 is %T, ok=%t\n", v2, v2, ok) // v2=, the type of v2 is string, ok=false
	//v3 := i.(int64) // 使用这个方式，如果i的类型不是int64，就会panic
	//fmt.Printf("v3=%d, the type of v3 is %T\n", v3, v3) // v3=13, the type of v3 is int64
	//v4 := i.([]int) // panic: interface conversion: interface {} is int64, not []int
	//fmt.Printf("the type of v4 is %T\n", v4)
}
