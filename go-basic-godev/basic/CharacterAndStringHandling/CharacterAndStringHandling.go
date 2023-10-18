package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "世界首发, 就在这里!"
	fmt.Println(len(s))
	fmt.Printf("%x\n", []byte(s))
	for _, b := range []byte(s) { // utf8编码 中文3字节
		fmt.Printf("%x ", b) // utf-8编码
	}
	fmt.Println()

	// rune 将 string进行utf-8解码，转换成 unicode
	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d, %x) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// 将s转换为Slice
	bytes := []byte(s)
	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		fmt.Printf("%c %v\n", r, size)
		bytes = bytes[size:]
	}

	// 直接转成rune，就是一个字符一个字符
	// rune是Decode之后的结果
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}

}
