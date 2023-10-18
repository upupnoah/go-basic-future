package main

import (
	"fmt"
	"net/http"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome Gopher")
	if err != nil {
		_ = fmt.Errorf("error: %v", err)
		return
	}
}

/*
个人理解：
1. 为何需要强转
	在函数签名一样的情况下， 可以这个函数签名定义成一个类型， 然后给这个类型实现具体的方法，那么强转过去的函数，
	就可以直接使用这些方法了，而不需要自己再次实现
2. 什么样的条件下可以进行强转？
	函数签名一样的情况下就可以强转
*/

func main() {
	// 因为他们的函数签名一样， 所以可以强转
	err := http.ListenAndServe(":8080", http.HandlerFunc(greeting) /*将greeting类型强转成HandlerFunc*/)
	if err != nil {
		_ = fmt.Errorf("error: %v", err)
	}
}
