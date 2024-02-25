package main

import "github.com/gin-gonic/gin"

func init() {
	gin.ForceConsoleColor()

}

func main() {
	server := gin.Default()
	server.GET("/hello", hello)
	server.Run(":8999")
}

func hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello",
	})
}
