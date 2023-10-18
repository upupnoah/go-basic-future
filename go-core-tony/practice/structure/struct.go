package main

import (
	"bytes"
	"fmt"
	"unsafe"
)

type Person struct {
	Name  string
	Phone string
	Addr  string
}

// Book Type embedding
type Book struct {
	Title string
	Person
}

type testZero struct {
	buf []byte // contents are the bytes buf[off : len(buf)]
	off int    // read at &buf[off], write at &buf[len(buf)]
}

func main() {
	var book Book
	book.Title = "Go language"
	book.Name = "NoahX"
	book.Phone = "10086"
	book.Addr = "HangZhou"
	fmt.Println(book)
	//fmt.Println(book.Author.Name)
	fmt.Println(book.Name)

	var b bytes.Buffer
	fmt.Println(unsafe.Sizeof(b))

	b.Write([]byte("Hello, Go"))
	fmt.Println(b.String())
	t := bytes.Buffer{}
	fmt.Println(unsafe.Sizeof(t))

	// 零值可用
	var t1 testZero
	k, v := t1.Write2()
	fmt.Println(k, v)

}

func (b *testZero) Write2() (n int, err error) {
	return 1, nil
}
