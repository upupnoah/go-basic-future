package main

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository/dao"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/service"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/web"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/web/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDB()
	u := initUser(db)
	server := initWebServer()

	// 注册路由
	u.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}

// initDB 初始化数据库
func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		// 只在初始化过程中 panic， panic 相当于整个 goroutine 结束
		// 引入了日志库之后， 会改成 log.Fatal 之类的
		panic(err)
	}
	// 初始化表(建表：检测表是否存在。不存在则建表，存在则根据 struct 的变化修改表)
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}

// initUser 初始化用户模块
func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDao(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

// initWebServer 初始化 web 服务
func initWebServer() *gin.Engine {
	server := gin.Default()

	// u.RegisterRoutesV1(server.Group("/users"))

	server.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"https://foo.com"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		ExposeHeaders:   []string{"x-jwt-token"}, // 允许前端获取这个响应头
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost:3000") {
				return true
			}
			if strings.HasPrefix(origin, "http://localhost:3001") { // node 的端口, 这里使用vscode端口转发了,因此还是localhost
				return true
			}
			// 请求的域名, 此时我mac上访问的是: http://192.168.31.38:3001
			// if strings.HasPrefix(origin, "http://192.168.31.38:3001") { // node 的端口, 这里使用vscode端口转发了,因此还是localhost
			// 	return true
			// }
			return strings.Contains(origin, "your custom domain...")
		},
		MaxAge: 12 * time.Hour,
	}))

	// 基于 cookie 的 session，不安全(浏览器能看到的你 session 的内容)
	// store := cookie.NewStore([]byte("secret"))

	// 基于 memstore 的 session，适用于单机
	// store := memstore.NewStore([]byte("oiX8LrevZyTBMNQcXS49djcmHVMmxggm"),
	// 	[]byte("N14EsDR03ubrCCYzHQvIPleU4rli8VzA"))

	// 基于 redis 的 session，适用于分布式
	// 参数说明：
	// 第一个参数：最大的连接数
	// 第二个参数：网络类型
	// 第三个参数：redis 地址
	// 第四个参数：密码
	// 第五个参数：加密用的 key
	store, err := redis.NewStore(16, "tcp", "localhost:6379", "", []byte("oiX8LrevZyTBMNQcXS49djcmHVMmxggm"),
		[]byte("N14EsDR03ubrCCYzHQvIPleU4rli8VzA"))
	if err != nil {
		panic(err)
	}

	// 设置 session 的选项
	store.Options(sessions.Options{
		// SameSite: http.SameSiteNoneMode,
		// Secure: true,
		// HttpOnly: true,
		MaxAge: 30,
	})

	// session 中间件
	server.Use(sessions.Sessions("mysession", store))

	// 登录校验中间件
	// server.Use(middleware.
	// 	NewLoginMiddlewareBuilder().CheckLogin())
	server.Use(middleware.NewLoginJWTMiddlewareBuilder().CheckLogin())
	return server
}
