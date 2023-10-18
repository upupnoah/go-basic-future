package main

import "net/http"

// 采用最简项目布局

func main() {
	// HandleFunc 的第二个参数
	// w 是用来操作返回给客户端来应答的
	// r 代表来自客户端的http请求
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world\n"))
		w.Write([]byte("upupqi"))
	})
	http.ListenAndServe(":8080", nil)
}
