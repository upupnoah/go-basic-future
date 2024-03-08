package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/domain"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository/cache"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository/dao"
)

var (
	ErrUserDuplicateUser = dao.ErrUserDuplicate
	ErrUserNotFound      = dao.ErrDataNotFound
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	FindById(ctx context.Context, id int64) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindByPhone(ctx context.Context, phone string) (domain.User, error)
}

type CachedUserRepository struct {
	dao   dao.UserDAO
	cache cache.UserCache
}

func NewUserRepository(dao dao.UserDAO, cache cache.UserCache) UserRepository {
	return &CachedUserRepository{
		dao:   dao,
		cache: cache,
	}
}

func (repo *CachedUserRepository) toDomain(u dao.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Email:    u.Email.String,
		Phone:    u.Phone.String,
		Password: u.Password,
		AboutMe:  u.AboutMe,
		Nickname: u.Nickname,
		Birthday: time.UnixMilli(u.Birthday),
	}
}

func (repo *CachedUserRepository) toEntity(u domain.User) dao.User {
	return dao.User{
		Id: u.Id,
		Email: sql.NullString{
			String: u.Email,
			Valid:  u.Email != "",
		},
		Phone: sql.NullString{
			String: u.Phone,
			Valid:  u.Phone != "",
		},
		Password: u.Password,
		Birthday: u.Birthday.UnixMilli(),
		AboutMe:  u.AboutMe,
		Nickname: u.Nickname,
	}
}

func (ur *CachedUserRepository) Create(ctx context.Context, user domain.User) error {
	err := ur.dao.Insert(ctx, ur.toEntity(user))
	return err
}

func (ur *CachedUserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
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
	user = ur.toDomain(u)

	// if set failed, just ignore (commonly, cache is not critical)
	// advanced method: consider the redis is down, keep the db alive
	// we can use rate limit to avoid the db being down

	// redis 如果出问题, 那么压力就会来到数据库上, 这时候可以使用限流来避免数据库被压垮

	// set cache
	err = ur.cache.Set(ctx, user)
	if err != nil {
		log.Println(err)
	}
	return user, nil
}

func (ur *CachedUserRepository) FindByIdV1(ctx context.Context, id int64) (domain.User, error) {
	user, err := ur.cache.Get(ctx, id)
	switch err {
	case nil:
		return user, err
	case cache.ErrCacheMiss:
		daoUser, err := ur.dao.FindById(ctx, id)
		if err != nil {
			return domain.User{}, err
		}
		user = ur.toDomain(daoUser)
		_ = ur.cache.Set(ctx, user)
		return user, nil
	default:
		return domain.User{}, err
	}
}

func (ur *CachedUserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := ur.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return ur.toDomain(u), nil
}

func (ur *CachedUserRepository) FindByPhone(ctx context.Context, phone string) (domain.User, error) {
	u, err := ur.dao.FindByPhone(ctx, phone)
	if err != nil {
		return domain.User{}, err
	}
	return ur.toDomain(u), nil
}
