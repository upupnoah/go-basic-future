package main

import (
	"encoding/gob"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	gob.Register(time.Time{}) // 用于 session 的序列化
	gin.ForceConsoleColor()   // 强制控制台颜色
}

func main() {
	server := InitWebServer()
	if err := server.Run(":8099"); err != nil {
		panic(err)
	}
}
