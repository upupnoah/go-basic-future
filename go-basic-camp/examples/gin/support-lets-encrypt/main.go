package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

// example for 1-line LetsEncrypt HTTPS servers.
func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
}

// example for custom autocert manager.
// func main() {
// 	r := gin.Default()

// 	// Ping handler
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.String(200, "pong")
// 	})

// 	m := autocert.Manager{
// 		Prompt:     autocert.AcceptTOS,
// 		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
// 		Cache:      autocert.DirCache("/var/www/.cache"),
// 	}

// 	log.Fatal(autotls.RunWithManager(r, &m))
// }
