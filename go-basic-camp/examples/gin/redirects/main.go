package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// web page redirect
	r.GET("/test1", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})

	r.POST("/test2", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/foo")
	})

	// router redirect
	r.GET("/test3", func(c *gin.Context) {
		c.Request.URL.Path = "/test4"
		r.HandleContext(c)
	})
	r.GET("/test4", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	r.Run(":8085")
}
