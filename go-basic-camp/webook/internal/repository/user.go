package repository

import (
	"context"
	"log"

	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/domain"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository/cache"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository/dao"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound       = dao.ErrDataNotFound
)

type UserRepository struct {
	dao   *dao.UserDAO
	cache *cache.UserCache
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (ur *UserRepository) Create(ctx context.Context, user domain.User) error {
	err := ur.dao.Insert(ctx, dao.User{
		Email:    user.Email,
		Password: user.Password,
	})
	return err
}

func (ur *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	// find first through cache
	user, err := ur.cache.Get(ctx, id)
	if err == nil {
		return user, nil
	}
	if err == cache.ErrCacheMiss {
		log.Println("cache miss")
	}
	// if not found in cache, find through db
	u, err := ur.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	user = domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}

	// if set failed, just ignore (commonly, cache is not critical)
	// advanced method: consider the redis is down, keep the db alive
	// we can use rate limit to avoid the db being down
	
	// redis 如果出问题, 那么压力就会来到数据库上, 这时候可以使用限流来避免数据库被压垮

	// set cache
	_ = ur.cache.Set(ctx, user)
	return user, nil
}

func (ur *UserRepository) FindByIdV1(ctx context.Context, id int64) (domain.User, error) {
	user, err := ur.cache.Get(ctx, id)
	switch err {
	case nil:
		return user, err
	case cache.ErrCacheMiss:
		daoUser, err := ur.dao.FindById(ctx, id)
		if err != nil {
			return domain.User{}, err
		}
		user = domain.User{
			Id:       daoUser.Id,
			Email:    daoUser.Email,
			Password: daoUser.Password,
		}
		_ = ur.cache.Set(ctx, user)
		return user, nil
	default:
		return domain.User{}, err
	}
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := ur.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}
