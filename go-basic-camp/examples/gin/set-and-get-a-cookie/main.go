package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		// 根据 cookie 名获取 cookie, 如果不存在则返回 error
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			// 设置 cookie
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run(":8081")
}
