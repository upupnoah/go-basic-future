package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrDataNotFound       = gorm.ErrRecordNotFound
	ErrUserDuplicateEmail = errors.New("user email duplicate")
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (ud *UserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.CreatedAt = now
	u.UpdatedAt = now
	err := ud.db.WithContext(ctx).Create(&u).Error
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		const uniqueIndexErrNo uint16 = 1062
		if mysqlError.Number == uniqueIndexErrNo {
			return ErrUserDuplicateEmail
		}
	}
	return err
}

func (ud *UserDAO) FindById(ctx context.Context, id int64) (User, error) {
	var u User
	err := ud.db.WithContext(ctx).First(&u, "id = ?", id).Error
	return u, err
}

func (ud *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	// 两种都可以, 第一种可读性好一些, 第二种简洁
	// err := ud.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	err := ud.db.WithContext(ctx).First(&u, "email = ?", email).Error
	return u, err
}

type User struct {
	Id       int64  `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"unique"`
	Password string

	CreatedAt int64 // 统一 UTC +0, 涉及到时间的时候, 再处理时区(转换)
	UpdatedAt int64
}
