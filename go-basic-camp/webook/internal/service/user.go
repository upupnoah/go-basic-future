package service

import (
	"context"

	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/domain"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository"
)

// UserService 用户服务
type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository)  *UserService{
	return &UserService {
		repo: repo,
	}
}

// SignUp 注册， 不是 Create ，因为有可能是邀请注册的，那就需要使用另一套逻辑
func (usvc *UserService) SignUp(ctx context.Context, u domain.User) error { // 不知道返回啥的时候就返回error
	// todo 考虑加密放在哪里
	//  存储到数据库
	return usvc.repo.Create(ctx, u)
}
