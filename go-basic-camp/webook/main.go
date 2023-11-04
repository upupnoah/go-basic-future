package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/web/user"
	"strings"
	"time"
)

func main() {
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"https://foo.com"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost:3000") {
				return true
			}
			return strings.Contains(origin, "your custom domain...")
		},
		MaxAge: 12 * time.Hour,
	}))

	u := user.NewHandler()
	u.RegisterRoutesV1(engine.Group("/users"))

	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
