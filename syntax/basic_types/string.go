package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 反引号：原生字符串，不会转义（可以多行，但是不能用\` 转义 ‘`
	var str1 string = `我可以换行
	这是新的行
	但是这里不能有反引号`

	var str2 string = "this is a raw string \\n" // 通过 \ 转义
	println(str1)
	println(str2)

	// 字符串拼接
	fmt.Println("Hello, " + "world!")

	// strings 包里面放着各种字符串相关操作的方法，需要的时候再查阅

	// 中文的场景
	fmt.Println(len("你好"))                         // 6
	fmt.Println(utf8.RuneCountInString("你好"))      // 2
	fmt.Println(utf8.RuneCountInString("你好hello")) // 7

	var s1 string = "hello"
	//s[0] = 'k' // ❌字符串的内容是不可变的
	s1 = "gopher" // 重新赋值是ok的
	fmt.Println(s1)

	var s2 = "中国人"
	fmt.Printf("the length of s = %d\n", len(s2)) //9
	for i := 0; i < len(s2); i++ {
		fmt.Printf("0x%x ", s2[i])
	}
	fmt.Println()

	var s3 = "中国人"
	fmt.Println("the character count in s is", utf8.RuneCountInString(s3)) // 3
	for _, c := range s3 {                                                 // range string 得到的 c 是 rune 类型
		fmt.Printf("0x%x ", c) // 0x4e2d 0x56fd 0x4eba
	}
	fmt.Println()

}
