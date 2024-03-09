package ioc

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/web"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/web/middleware"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/pkg/ginx/middleware/ratelimit"
)

func InitWebServer(middlewares []gin.HandlerFunc, userHandler web.UserHandler) *gin.Engine {
	server := gin.Default()
	server.Use(middlewares...)
	userHandler.RegisterRoutes(server)
	return server
}

func InitGinMiddlewares(redisClient redis.Cmdable) []gin.HandlerFunc {
	ljmb := middleware.NewLoginJWTMiddlewareBuilder()
	return []gin.HandlerFunc{
		cors.New(cors.Config{
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
		}),

		ratelimit.NewBuilder(redisClient, time.Second, 100).Build(),

		ljmb.SetIgnorePaths("/users/login",
			"/users/signup", "/users/login_sms/code/send", "/users/login_sms").CheckLoginJWT(),
	}
}
