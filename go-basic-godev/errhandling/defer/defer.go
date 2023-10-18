package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/upupnoah/go-basic-future/go-basic-godev/errhandling/fib"
)

func tryDefer() {
	// defer的好处是不怕中间return
	// 后进先出
	defer fmt.Println(9999)
	defer fmt.Println(999)
	defer fmt.Println(99)
	defer fmt.Println(9)
	fmt.Println(2)
	fmt.Println(3)
	//panic("error occurred")
	return
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		//panic(err)
		/* 正确的错误处理 */
		fmt.Println("Error: ", err)
		return
	}

	err = errors.New("this is a custom error")

	// 知道具体错误， 就具体处理
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err) // 如果不是一个pathError， 那么就会 panic
		} else {
			fmt.Printf("%s %s %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		//fmt.Println(f())
		fmt.Fprintln(writer, f())
	}

}

func main() {
	tryDefer()
	writeFile("Go-Development-Engineer/errhandling/defer/fib.txt")
}
