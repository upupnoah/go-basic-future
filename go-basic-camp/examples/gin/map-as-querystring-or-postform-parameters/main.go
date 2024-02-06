package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
// Content-Type: application/x-www-form-urlencoded

// names[first]=thinkerou&names[second]=tianou

// curl -X "POST" "http://localhost:8081/post?ids%5Ba%5D=1234&ids%5Bb%5D=hello" \
//      -H 'Content-Type: application/x-www-form-urlencoded; charset=utf-8' \
//      --data-urlencode "names[first]=thinkerou" \
//      --data-urlencode "names[second]=tianou"

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		log.Printf("ids: %v; names: %v", ids, names)
	})
	router.Run(":8081")
}
