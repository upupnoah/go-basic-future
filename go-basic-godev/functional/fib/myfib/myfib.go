package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// 函数， 返回值为 func() int
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 既然函数是一等公民，可以作为类型， 变量
// 那我们让这个类型（func int）实现一下 Reader接口
// 要是实现Reader接口就要让这个类型中 实现Read(p []byte) (int, error)
// 给函数实现方法， 则需要先定义一个类型
type intGen func() int // 这个操作是给函数 定义一个类型

// 实现一个类型的方法 -> 让这个类型作为接收者

/*
	说是给函数类型(func() int) 实现Reader接口, 实际上就是 给这个类型 定义一个方法
	也就是 实现 io.Reader接口中声明的方法
*/

func (f intGen) Read(p []byte) (n int, err error) {
	//panic("implement me")
	next := f() // 读取下一个
	// 因为f是一个生成器， 可以无限读取下一个, 因此需要做一个限制
	if next > 10000 { // Read 返回 io.EOF -> 文件读到头了
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func main() {
	// 在go语言中， 函数是一等公民， 因此可以作为变量
	// 这里是生成一个生成器 f为一个func int类型的变量
	f := fibonacci()
	//fmt.Printf("%T %v\n", f, f)
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//var reader io.Reader = f
	printFileContents(f)

}
