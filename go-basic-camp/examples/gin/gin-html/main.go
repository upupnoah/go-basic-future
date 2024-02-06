package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/render-html", renderHtml)

	r.Run(":8080")
}

func renderHtml(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
