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
	// ErrUserNotFound       = repository.ErrUserNotFound
	ErrInvalidUserOrPassword = errors.New("账号/邮箱 或 密码错误")
	ErrUserNotFound          = errors.New("用户不存在")
)

// UserService 用户服务
type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// SignUp 注册， 不是 Create ，因为有可能是邀请注册的，那就需要使用另一套逻辑
func (usvc *UserService) SignUp(ctx context.Context, u domain.User) error { // 不知道返回啥的时候就返回error
	// 密码加密
	bPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bPassword)

	//  存储到数据库
	return usvc.repo.Create(ctx, u)
}

func (usvc *UserService) Login(ctx context.Context, email, password string) (domain.User, error) {
	// 先找用户
	u, err := usvc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrUserNotFound
	}
	if err != nil {
		return domain.User{}, err
	}
	// 比较密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		// DEBUG 日志，方便测试
		// log.Printf("password compare failed, err: %v", err)
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return domain.User{}, nil
}

func (usvc *UserService) UpdateProfile(ctx context.Context, u domain.User) error {
	return usvc.repo.Update(ctx, u)
}