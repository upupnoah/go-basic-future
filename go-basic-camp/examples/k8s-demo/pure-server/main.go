package main

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/upupnoah/go-basic-future/go-basic-camp/examples/k8s-demo/pure-server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	gin.ForceConsoleColor()
}

func main() {
	server := initWebServer()
	server.GET("/hello", hello)
	initDB()
	initRedis()
	server.Run(":8999")
}

func hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello",
	})
}

func initRedis() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Addr,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DB,
	})
	pong, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	} else {
		log.Println("Redis connected:", pong)
	}
}

func initDB() {
	config := config.Config
	db, err := gorm.Open(mysql.Open(config.DB.DSN))
	if err != nil {
		panic(err)
	}

	err = initTables(db)
	if err != nil {
		panic(err)
	}

	// show current database name
	currentDatabase := db.Migrator().CurrentDatabase()
	log.Println("currentDatabase", currentDatabase)
}

func initWebServer() *gin.Engine {
	srv := gin.Default()
	// 解决 [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	err := srv.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}
	srv.Use(cors.New(cors.Config{
		// AllowOrigins: []string{"*"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"x-jwt-token"}, // 允许前端获取这个响应头
		AllowOriginFunc: func(origin string) bool {
			// 只允许来自 noah.webook.com 和 localhost 的请求
			return origin == "https://noah.webook.com" || strings.HasPrefix(origin, "http://localhost")
		},
		MaxAge: 12 * time.Hour,
	}))

	return srv
}

type User struct {
	Id       int64  `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"unique"`
	Password string

	CreatedAt int64 // 统一 UTC +0, 涉及到时间的时候, 再处理时区(转换)
	UpdatedAt int64
}

func initTables(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
