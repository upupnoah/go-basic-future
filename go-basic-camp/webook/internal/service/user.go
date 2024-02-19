package service

import (
	"context"
	"errors"

	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/domain"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail

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
