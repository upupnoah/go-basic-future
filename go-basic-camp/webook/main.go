package main

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository/dao"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/service"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
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
	
	ud := dao.NewUserDao(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	server := gin.Default()

	// u.RegisterRoutesV1(server.Group("/users"))

	server.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"https://foo.com"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
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

	// 注册路由
	u.RegisterRoutes(server)
	err = server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
