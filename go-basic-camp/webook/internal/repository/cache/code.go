package cache

import (
	"context"
	_ "embed"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	//go:embed lua/set_code.lua
	luaSetCode string
	//go:embed lua/verify_code.lua
	luaVerifyCode string
	// 验证码发送频繁
	ErrCodeSendTooMany = errors.New("code send too many")
	// 发送验证码遇到未知错误
	ErrUnknownForCode         = errors.New("unknown error for code")
	ErrCodeVerifyTooManyTimes = errors.New("code verify too many times")
)

type CodeCache struct {
	client redis.Cmdable
}

func NewCodeCache(client redis.Cmdable) *CodeCache {
	return &CodeCache{
		client: client,
	}
}

func (cc *CodeCache) Set(ctx context.Context, biz, phone, code string) error {
	key := cc.key(biz, phone)
	res, err := cc.client.Eval(ctx, luaSetCode, []string{key}, code).Int()
	if err != nil {
		return err
	}
	switch res {
	case 0:
		return nil
	case -1:
		// 最近发过验证码
		return ErrCodeSendTooMany
	default:
		// 系统错误, 比如-2, key 冲突
		// 其他响应码, unknown error
		// TODO: log
		return ErrUnknownForCode
	}
}

func (cc *CodeCache) Verify(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	res, err := cc.client.Eval(ctx, luaVerifyCode, []string{cc.key(biz, phone)}, inputCode).Int()
	if err != nil {
		return false, err
	}
	switch res {
	case 0:
		return true, nil
	case -1:
		// 验证次数耗尽
		return false, ErrCodeVerifyTooManyTimes
	default:
		// 验证码不对
		return false, nil
	}
}

func (cc *CodeCache) key(biz, phone string) string {
	return fmt.Sprintf("phone_code:%s:%s", biz, phone)
}
