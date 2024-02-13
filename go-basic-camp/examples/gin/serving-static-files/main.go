package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 这三个 method, 第一个参数都是请求路径, 第二个参数都是文件系统路径

	// http://localhost:8081/assets/js/test.js
	router.Static("/assets", "./assets")

	// http://localhost:8081/more_static/index.html
	router.StaticFS("/more_static", http.Dir("my_file_system"))

	// 这个方法用于提供单个静态文件。它映射一个具体的请求路径到一个具体的文件
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8081")
}
