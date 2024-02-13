package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Bind query string or post data

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", startPage)
	route.POST("/testing", startPage)
	route.Run(":8080")
}

// 总结
// Bind, BindJSON: 错误会返回 400
// 		BindJSON: 只能用于 JSON
// ShouldBind, ShouldBindJSON: 错误需要我们自己处理

func startPage(c *gin.Context) {
	var person Person

	// 请求 URL
	// Get 请求: http://localhost:8080/testing?name=Noah&address=%E6%B5%99%E6%B1%9F&birthday=2024-01-08
	// Post 请求: http://localhost:8080/testing

	// 手动处理 Get
	// person.Name = c.Query("name")
	// person.Address = c.Query("address")
	// person.Birthday, _ = time.Parse("2006-01-02", c.Query("birthday"))

	// 手动处理 Post
	// person.Name = c.PostForm("name")
	// person.Address = c.PostForm("address")
	// person.Birthday, _ = time.Parse("2006-01-02", c.PostForm("birthday"))

	// Bind(), 如果数据绑定失败, 他会返回 400 error
	// c.Bind(&person)
	// log.Println(person.Name)
	// log.Println(person.Address)
	// log.Println(person.Birthday)

	// ShouldBind(), 如果数据绑定失败, 需要我们自己处理 error
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48

	err := c.ShouldBindJSON(&person)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest) // 中断请求的处理流程(阻止后续所有处理器或中间件被调用)并设置响应状态码
		return
	}

	log.Println(person.Name)
	log.Println(person.Address)
	log.Println(person.Birthday)
	c.JSON(200, gin.H{"name": person.Name, "address": person.Address, "birthday": person.Birthday})
	// c.String(200, "Success")
}
