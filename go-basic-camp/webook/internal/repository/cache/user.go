package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/domain"
)

// var ErrKeyNotExist = redis.Nil
var ErrCacheMiss = errors.New("cache: miss")

type UserCache interface {
	Get(ctx context.Context, id int64) (domain.User, error)
	Set(ctx context.Context, user domain.User) error
}

type RedisUserCache struct {
	cmd        redis.Cmdable
	expiration time.Duration // cache expiration time
}

func NewUserCache(cmd redis.Cmdable) UserCache {
	return &RedisUserCache{
		cmd:        cmd,
		expiration: time.Minute * 15,
	}
}

func (uc *RedisUserCache) Get(ctx context.Context, id int64) (domain.User, error) {
	key := uc.key(id)
	data, err := uc.cmd.Get(ctx, key).Result()
	if err != nil {
		return domain.User{}, ErrCacheMiss
	}
	var user domain.User
	err = json.Unmarshal([]byte(data), &user)
	return user, err
}

func (uc *RedisUserCache) Set(ctx context.Context, user domain.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	key := uc.key(user.Id)
	return uc.cmd.Set(ctx, key, data, uc.expiration).Err()
}

func (uc *RedisUserCache) key(id int64) string {
	return fmt.Sprintf("user:info:%d", id)
}
