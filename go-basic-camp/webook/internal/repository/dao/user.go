package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/domain"
	"gorm.io/gorm"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
	ErrUserNotFound       = gorm.ErrRecordNotFound
	// ErrInvalidUserOrPassword = gorm.ErrRecordNotFound
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (dao *UserDao) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	// err := dao.db.WithContext(ctx).First(&u, "email=?", email).Error
	return u, err
}

func (dao *UserDao) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli() // 毫秒数, 根据具体业务场景
	u.Ctime = now
	u.Utime = now

	// 检查邮箱是否冲突, 与底层强耦合
	err := dao.db.WithContext(ctx).Create(&u).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok { // 因为 email 字段设置了 unique 约束（唯一索引），所以会报错
		if mysqlErr.Number == 1062 {
			return ErrUserDuplicateEmail
		}
	}
	return err
	// return dao.db.WithContext(ctx).Create(&u).Error
}

func (dao *UserDao) Update(ctx context.Context, u domain.User) error {
	// 更新 u 中的信息
	return dao.db.WithContext(ctx).Save(&u).Error
}

// User 直接对应数据库 User 表结构， 还有其他叫法：entity， model，PO(persistent object)
type User struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string
	// 创建时间, 毫秒数， 不适用 time.Time 是因为这个类型和时区有关，不利于跨时区
	Ctime int64
	// 更新时间, 毫秒数
	Utime int64
}
