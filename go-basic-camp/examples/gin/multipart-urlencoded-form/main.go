package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/form_post", func(ctx *gin.Context) {
		message := ctx.PostForm("message")
		nick := ctx.DefaultPostForm("nick", "anonymous")
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8081")
}
