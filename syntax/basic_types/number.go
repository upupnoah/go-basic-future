package main

import (
	"fmt"
	"math"
)

// 全局不适用
// d := 4 ❌
// var d int = 4 ✅

func main() {
	// 定义变量的三种方式
	var a int
	var b int = 2
	c := 3 // 这种只能在 func 中使用 （不能在全局使用）
	fmt.Println("int: ", a, b, c)

	var f1 float32 = 1.32
	var f2 = 1.64 // 1.64 默认是 float64，因此可以省略
	fmt.Printf("float32: %f\nfloat64: %f\n", f1, f2)

	var c1 complex64 = complex(1, 2)
	c2 := 3 + 4i
	fmt.Printf("complex64: %.1f\ncomplex128: %.1f\n", real(c1), imag(c1)) // 区分实部 和 虚部
	fmt.Printf("%+v\n", c2)                                               // 输出完整的复数

	// 极值
	Extremum()

	// 类型转换
	TypeConversion()
}

// Extremum 极值
func Extremum() {
	// int族 和 uint族 都有最大值最小值

	// float族 没有 最小值， 只有 最大值 和 最小正数
	fmt.Println("Maximum value of float64:", math.MaxFloat64)
	fmt.Println("Smallest nonzero float64:", math.SmallestNonzeroFloat64)

	// float32同理
}

// TypeConversion 类型转换
func TypeConversion() {
	// Go 语言不支持隐式类型转换，也就是说在 Go 语言中，即使是从窄类型到宽类型的转换也必须显式声明
	var a int = 1
	//var b float64 = a // ❌ cannot use a (type int) as type float64 in assignment
	var b float64 = float64(a)
	fmt.Println("a:", a, "b:", b)

	// 但是对于常量的转换，可以不用显式声明
	const c = 5
	var d float64 = c
	fmt.Println("c:", c, "d:", d)
}
