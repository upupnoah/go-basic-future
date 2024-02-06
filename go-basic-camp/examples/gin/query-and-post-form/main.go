package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// POST /post?id=1234&page=1
func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		log.Printf("id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)
	})
	router.Run(":8085")
}
