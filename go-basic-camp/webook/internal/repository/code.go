package repository

import (
	"context"

	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository/cache"
)

// 复用 cache 包中的错误
var (
	ErrCodeVerifyTooManyTimes = cache.ErrCodeVerifyTooManyTimes
	ErrCodeSendTooMany        = cache.ErrCodeSendTooMany
)

type CodeRepository interface {
	Store(ctx context.Context, biz, phone, code string) error
	Verify(ctx context.Context, biz, phone, inputCode string) (bool, error)
}

type CachedCodeRepository struct {
	cache cache.CodeCache
}

func NewCodeRepository(cache cache.CodeCache) CodeRepository {
	return &CachedCodeRepository{
		cache: cache,
	}
}

func (cr *CachedCodeRepository) Store(ctx context.Context, biz, phone, code string) error {
	err := cr.cache.Set(ctx, biz, phone, code)
	return err
}

func (cr *CachedCodeRepository) Verify(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	res, err := cr.cache.Verify(ctx, biz, phone, inputCode)
	if err != nil {
		return false, err
	}
	return res, nil
}
