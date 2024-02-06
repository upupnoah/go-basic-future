package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl -X POST http://localhost:8080/upload \
//   -F "upload[]=@/Users/appleboy/test1.zip" \
//   -F "upload[]=@/Users/appleboy/test2.zip" \
//   -H "Content-Type: multipart/form-data"

// Multiple files
func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// Upload the file to specific dst.
			c.SaveUploadedFile(file, "./files/"+file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!\n", len(files)))
	})
	router.Run(":8081")
}

// curl -X POST http://localhost:8080/upload \
//   -F "file=@/Users/appleboy/test.zip" \
//   -H "Content-Type: multipart/form-data"

// Single file
// func main() {
// 	router := gin.Default()
// 	// Set a lower memory limit for multipart forms (default is 32 MiB)
// 	router.MaxMultipartMemory = 8 << 20 // 8 MiB
// 	router.POST("/upload", func(c *gin.Context) {
// 		// single file
// 		file, _ := c.FormFile("file")
// 		log.Println(file.Filename)

// 		// Upload the file to specific dst.
// 		c.SaveUploadedFile(file, dst)

// 		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
// 	})
// 	router.Run(":8080")
// }
