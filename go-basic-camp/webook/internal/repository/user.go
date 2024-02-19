package repository

import (
	"context"

	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/domain"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository/dao"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound       = dao.ErrDataNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
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
	u, err := ur.dao.FindById(ctx, id)
	return domain.User{
		Email:    u.Email,
		Password: u.Password,
	}, err
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
