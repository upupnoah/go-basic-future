package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 函数外部初始化变量必须有var 关键字或者其他关键字
//var aa = 3
//var ss = "kkk"
/* 简单写法 */
var (
	aa = 3
	ss = "kkk"
	bb = true
)

// 变量默认值
func variableZeroValue() {
	var a int // 这么定义符合正常思维模式， 先想到名字， 再想到类型
	var s string
	//fmt.Println(a, s)
	//%q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	fmt.Printf("%d %q\n", a, s)
}

// 变量初始化
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// 类型自动推断
func variableTypeExtrapolate() {
	var a, b, c, s = 3, 4, true, "Extrapolate"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	// var能不用就不用
	a, b, c, s := 3, 4, true, "Extrapolate" //第一次需要使用 ":"
	b = 5                                   // 以后就不能再用了
	fmt.Println(a, b, c, s)
}

func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	// 欧拉公式
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

// func triangle() {
// 	var a, b int = 3, 4
// 	var c int = int(math.Sqrt(float64(a*a + b*b)))
// 	fmt.Println(c)
// }

// go没有隐式类型转换
func triangle() {
	var a, b int = 3, 4
	var c int = calcTriangle(a, b)
	fmt.Println(c)
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

// 常量
// 如果没有声明const的类型， 那么就是文本， 使用的时候就不用强制转换了
func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

// 枚举
func enums() {
	const (
		// cpp = 0
		// java = 1
		// python = 2
		// Go-Development-Engineer = 3
		cpp = iota
		_
		java
		python
		golang
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func MyPow(a *int) {
	*a = *a * *a
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeExtrapolate()
	variableShorter()
	fmt.Println(aa, bb, ss)
	euler()
	triangle()
	consts()
	enums()
	fmt.Printf("%T\n", ss)
	abc := 123

	fmt.Printf("%T\n", abc)
}
