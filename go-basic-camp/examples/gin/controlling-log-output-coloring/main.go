package main

import "github.com/gin-gonic/gin"

func main() {
	// Force log's color
	gin.ForceConsoleColor()

	// Disable log's color
	// gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
