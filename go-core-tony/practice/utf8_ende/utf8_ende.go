package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

// rune -> []byte
func encodeRune() {
	var r rune = 0x4E2D
	fmt.Printf("the unicode character is %c\n", r) // 中
	buf := make([]byte, 3)
	_ = utf8.EncodeRune(buf, r)                       // 对rune进行utf-8编码
	fmt.Printf("utf-8 representation is 0x%X\n", buf) // 0xE4B8AD
}

// []byte -> rune
func decodeRune() {
	var buf = []byte{0xE4, 0xB8, 0xAD}
	r, _ := utf8.DecodeRune(buf)                                                             // 对buf进行utf-8解码
	fmt.Printf("the unicode character fater decoding [0xE4, 0xB8, 0xAD] is %s\n", string(r)) // 中
}

func dumpBytesArray(arr []byte) {
	fmt.Printf("[")
	for _, b := range arr {
		fmt.Printf("%c ", b)
	}
	fmt.Printf("]\n")
}

func main() {
	var s = "hello"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // 将string类型变量地址显示转型为reflect.StringHeader
	fmt.Printf("0x%x\n", hdr.Data)
	p := (*[5]byte)(unsafe.Pointer(hdr.Data)) // 获取Data字段所指向的数组的指针
	dumpBytesArray((*p)[:])                   // [h e l l o] 输出底层数组的内容
}
