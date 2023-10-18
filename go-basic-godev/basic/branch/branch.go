package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	// switch 里面可以没有表达式， 直接在case中写上条件即可
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score)) // panic会中断程序的执行，并且会报错
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
		// default:
		// 	panic(
		// 		fmt.Sprintf(
		// 			"Wrong score: %d", score))
	}
	return g
}

func main() {
	const filename = "abcd.txt"
	// if else
	// contents, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }

	// 可以直接写在一起
	if contents, err := ioutil.ReadFile(filename); err == nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	// 已经出了作用域了
	// fmt.Println(contents)
	fmt.Printf("%s\n", grade(100))
}
