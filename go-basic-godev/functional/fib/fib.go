package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 定义一个生成器
//func fibonacci() func() int {
//	a, b := 0, 1
//	return func() int {
//		a, b = b, a+b // 每次迭代 0 1 1 2 3 5 8
//		return a
//	}
//}

// 这个生成器返回一个 func() int
// 我们给 fibonacci() 这个生成器实现一个 io.Reader, 那么同样也可以打印了
//func fibonacci() func() int {
//	a, b := 0, 1
//	return func() int {
//		a, b = b, a+b
//		return a // 对于这个闭包， 返回一个int
//	}
//}
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

// 既然函数能够作为变量， 作为参数， 那么也能作为接收者， 接收者只是一个特殊的参数而已
// 只要是一个类型， 就可以实现Reader接口（实现 Read method）， 这就是go语言中灵活的地方
func (g intGen) Read(p []byte) (n int, err error) {
	/*
		调用g() 就取得了下一个元素
		又因为这是一个fibonacci生成器， 永远来读不完
		所以可以设置一个限定调试
	*/
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	// 这个操作比较底层， 我们让实现了 Reader 的人帮我们代理一下
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	//fmt.Println(f()) //从 fibonacci(1) 开始
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	// 既然实现了 Reader接口， 那么就可以当文件来用
	//printFileContents(f)
	printFileContents(f)
}
