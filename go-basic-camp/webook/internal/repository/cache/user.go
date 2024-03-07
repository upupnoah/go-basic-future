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

type UserCache struct {
	cmd        redis.Cmdable
	expiration time.Duration // cache expiration time
}

func NewUserCache(cmd redis.Cmdable) *UserCache {
	return &UserCache{
		cmd:        cmd,
		expiration: time.Minute * 15,
	}
}

// 
func (uc *UserCache) Get(ctx context.Context, id int64) (domain.User, error) {
	key := uc.key(id)
	data, err := uc.cmd.Get(ctx, key).Result()
	if err != nil {
		return domain.User{}, ErrCacheMiss
	}
	var user domain.User
	err = json.Unmarshal([]byte(data), &user)
	return user, err
}

func (uc *UserCache) Set(ctx context.Context, user domain.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	key := uc.key(user.Id)
	return uc.cmd.Set(ctx, key, data, uc.expiration).Err()
}

func (uc *UserCache) key(id int64) string {
	return fmt.Sprintf("user:info:%d", id)
}
