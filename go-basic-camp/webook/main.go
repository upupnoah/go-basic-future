package main

import (
	"encoding/gob"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	redisV9 "github.com/redis/go-redis/v9"

	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/config"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository/dao"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/service"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/web"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/web/middleware"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/pkg/ginx/middleware/ratelimit"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	gob.Register(time.Time{}) // 用于 session 的序列化
	gin.ForceConsoleColor()   // 强制控制台颜色
}

func main() {
	db := initDB()
	srv := initWebServer()
	initUser(srv, db)
	err := srv.Run(":8099")
	if err != nil {
		panic(err)
	}
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	if err != nil {
		panic(err)
	}

	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}

	return db
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
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			if strings.HasPrefix(origin, "http://noah.webook.com") {
				return true
			}
			return strings.Contains(origin, "your custom domain...")
		},
		MaxAge: 12 * time.Hour,
	}))
	redisClient := redisV9.NewClient(&redisV9.Options{
		Addr: config.Config.Redis.Addr,
	})
	// 限流: 1s 100次
	srv.Use(ratelimit.NewBuilder(redisClient, time.Second, 100).Build())

	// store := cookie.NewStore([]byte("secret"))

	// memstore
	// store := memstore.NewStore([]byte("TL0qRTxsVIfQ25l4TFMXj8XRfPwb6MgO"), []byte("Ojtmwv5Xbx1YP513SVnrpXx8wHyBhSa0"))
	// redis store
	store, err := redis.NewStore(16, "tcp", config.Config.Redis.Addr, config.Config.Redis.Password, []byte(""))
	if err != nil {
		panic(err)
	}
	srv.Use(sessions.Sessions("WebookSession", store))

	ljmb := middleware.NewLoginJWTMiddlewareBuilder()
	srv.Use(ljmb.SetIgnorePaths("/users/login", "/users/signup").CheckLoginJWT())
	return srv
}

func initUser(server *gin.Engine, db *gorm.DB) {
	ud := dao.NewUserDAO(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)
	u := web.NewUserHandler(us)
	u.RegisterRoutes(server)
}
