package dao

import "gorm.io/gorm"

func InitTables(db *gorm.DB) error {
	// 后续有其他表，可以在这里添加
	err := db.AutoMigrate(&User{})
	return err
}
