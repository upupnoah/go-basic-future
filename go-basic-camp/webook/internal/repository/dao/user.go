package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (dao *UserDao) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli() // 毫秒数, 根据具体业务场景
	u.Ctime = now
	u.Utime = now
	return dao.db.WithContext(ctx).Create(&u).Error
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
