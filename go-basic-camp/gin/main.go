package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	// 通过 engine.httpMethod() 方法注册路由
	// 静态路由
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	// 参数路由
	engine.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		//c.JSON(http.StatusOK, gin.H{
		//	"name": name,
		//})
		c.String(http.StatusOK, "参数路由: Hello %s", name)
	})

	// 通配符路由
	engine.GET("/views/*.html", func(c *gin.Context) {
		c.String(http.StatusOK, "通配符路由: %s", c.FullPath())
	})

	// 查询参数 and QueryMap
	engine.GET("/order", func(c *gin.Context) {
		//qid := c.Query("id")
		name := c.Query("name")
		mp := c.QueryMap("id")
		c.String(http.StatusOK, "name: %s\nid1: %s, id2: %s, id3: %s", name, mp["1"], mp["2"], mp["3"])
	})

	// 通过 go func() 启动一个协程
	//go func() {
	//	err := engine.Run(":8081")
	//	if err != nil {
	//		return
	//	}
	//}()
	err := engine.Run(":8080")

	if err != nil {
		return
	}
}
