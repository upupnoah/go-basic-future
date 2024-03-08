//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository/cache"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository/dao"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/service"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/web"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/ioc"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 第三方依赖
		ioc.InitDB, ioc.InitRedis,

		// dao
		dao.NewUserDAO,

		// cache
		cache.NewCodeCache, cache.NewUserCache,

		// repo
		repository.NewCodeRepository, repository.NewUserRepository,

		// service
		ioc.InitSMSService, service.NewUserService, service.NewCodeService,

		// handler
		web.NewUserHandler,

		// server
		ioc.InitGinMiddlewares, ioc.InitWebServer,
	)
	return gin.Default()
}
