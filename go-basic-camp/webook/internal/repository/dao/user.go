package dao

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrDataNotFound  = gorm.ErrRecordNotFound
	ErrUserDuplicate = errors.New("user email duplicate")
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
			// 邮箱冲突 or 手机号码冲突
			return ErrUserDuplicate
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

func (ud *UserDAO) FindByPhone(ctx context.Context, phone string) (User, error) {
	var u User
	err := ud.db.WithContext(ctx).First(&u, "phone = ?", phone).Error
	return u, err
}

type User struct {
	Id    int64          `gorm:"primaryKey;autoIncrement"`
	Email sql.NullString `gorm:"unique"`

	// 由于密码不需要设置成 unique(唯一索引), 因此他为""(空字符串)也没关系
	Password string
	Nickname string `gorm:"type=varchar(128)"`

	Birthday int64  // YYYY-MM-DD
	AboutMe  string `gorm:"type=varchar(4096)"`

	// 代表这是一个可以为 NULL 的列
	// 唯一索引允许有多个 NULL, 但是不允许多个 ""
	Phone     sql.NullString `gorm:"unique"`
	CreatedAt int64          // 统一 UTC +0, 涉及到时间的时候, 再处理时区(转换)
	UpdatedAt int64
}
