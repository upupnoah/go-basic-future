package gorm

import (
	// _ "example.com/my_mysql_driver"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectMySQL() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:13316)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func advancedConnectMySQL() {
	_, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@tcp(127.0.0.1:13316)/gorm_demo?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                              // default size for string fields
		DisableDatetimePrecision:  true,                                                                             // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                             // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                             // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                            // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

// 使用自定义 MySQL 驱动
// func customizeDriver() {
// 	_, err := gorm.Open(mysql.New(mysql.Config{
// 		DriverName: "my_mysql_driver",
// 		DSN:        "root:root@tcp(localhost:13316)/gorm_demo?charset=utf8&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
// 	}), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// }

// 使用现有数据库连接初始化 *gorm.DB
// func existDBConnection() {
// 	sqlDB, err := sql.Open("mysql", "mydb_dsn")
// 	_, err = gorm.Open(mysql.New(mysql.Config{
// 		Conn: sqlDB,
// 	}), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// }

// PostgreSQL
// func connectPostgreSQL() {
// 	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
// 	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	// 禁用启动准备语句缓存(prepared statement cache)
// 	_, err = gorm.Open(postgres.New(postgres.Config{
// 		DSN:                  "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
// 		PreferSimpleProtocol: true, // disables implicit prepared statement usage
// 	}), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// }

// connection pool
// func connectionPool() {
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
// 	sqlDB.SetMaxIdleConns(10)

// 	// SetMaxOpenConns sets the maximum number of open connections to the database.
// 	sqlDB.SetMaxOpenConns(100)

// 	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
// 	sqlDB.SetConnMaxLifetime(time.Hour)
// }
