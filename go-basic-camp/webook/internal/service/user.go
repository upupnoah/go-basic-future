package service

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/domain"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository"
)

var (
	ErrUserDuplicateUser = repository.ErrUserDuplicateUser

	// 不能直接告诉用户是密码错误还是邮箱错误, 否则会泄露用户信息
	ErrInvalidUserOrPassword = errors.New("invalid email or password")
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) SignUp(ctx context.Context, user domain.User) error {
	// 处理加密
	b, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(b)
	// 处理数据库
	err = us.repo.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) Login(ctx context.Context, user domain.User) (domain.User, error) {
	// find user by email
	u, err := us.repo.FindByEmail(ctx, user.Email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	// user exist, verify password
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, err
}

func (us *UserService) Profile(ctx context.Context, id int64) (domain.User, error) {
	return us.repo.FindById(ctx, id)
}

func (us *UserService) FindOrCreate(ctx context.Context, phone string) (domain.User, error) {
	// First, find user by phone, we believe that the majority of users are existing users.
	// 快路径
	user, err := us.repo.FindByPhone(ctx, phone)
	if err != repository.ErrUserNotFound {
		// 1. err == nil, u 可用
		// 2. err != nil, server error
		return user, err
	}
	// 慢路径
	// 用户没找到, 就通过手机号创建一个
	err = us.repo.Create(ctx, domain.User{
		Phone: phone,
	})
	if err != nil && err != repository.ErrUserDuplicateUser {
		return domain.User{}, err
	}
	// 要么 err == nil, 要么 ErrDuplicateUser, 也代表用户存在
	// 主从延迟, 理论上来讲, 强制走主库
	// 需要返回一个包含 uid 的 domainUser
	return us.repo.FindByPhone(ctx, phone)
}

func (us *UserService) FindById(ctx context.Context,
	uid int64) (domain.User, error) {
	return us.repo.FindById(ctx, uid)
}
