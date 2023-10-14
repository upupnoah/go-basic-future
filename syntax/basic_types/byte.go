package main

import "fmt"

func Byte() {
	var a byte = 'a'
	fmt.Println(a) // 97

	var str string = "this is string"
	var bs []byte = []byte(str)
	bs[0] = 'T' // 不会修改原字符串
	fmt.Printf("%v\n", bs)
}

func main() {
	Byte()
}
