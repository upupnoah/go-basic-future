package ioc

import (
	"github.com/redis/go-redis/v9"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/config"
)

func InitRedis() redis.Cmdable {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.Config.Redis.Addr,
	})
	return rdb
}
