// 控制反转
package ioc

import (
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/config"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	if err != nil {
		panic(err)
	}

	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}

	return db
}
