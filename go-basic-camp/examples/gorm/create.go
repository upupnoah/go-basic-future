package gorm

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// https://gorm.io/docs/create.html

type User struct {
	Name      string
	Age       int
	Birthday  time.Time
	ID        uint `gorm:"primaryKey"` // 设置主键
	Role      string
	UUID      uuid.UUID
	CreatedAt time.Time
	Active bool
}

func createRecord() {
	// db, err := gorm.Open(mysql.New(mysql.Config{
	// 	DSN:                       "root:root@tcp(127.0.0.1:13316)/gorm_demo?charset=utf8&parseTime=True&loc=Local", // data source name
	// 	DefaultStringSize:         256,                                                                              // default size for string fields
	// 	DisableDatetimePrecision:  true,                                                                             // disable datetime precision, which not supported before MySQL 5.6
	// 	DontSupportRenameIndex:    true,                                                                             // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
	// 	DontSupportRenameColumn:   true,                                                                             // `change` when rename column, rename column not supported before MySQL 8, MariaDB
	// 	SkipInitializeWithVersion: false,                                                                            // auto configure based on currently MySQL version
	// }), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	db := connectMySQL()

	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	result := db.Create(&user)
	fmt.Printf("result: %v\n", result)

	// create multiple records with Create()
	// users := []*User{
	// 	{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
	// 	{Name: "Jackson", Age: 19, Birthday: time.Now()},
	// }
	// result = db.Create(&users)
	// fmt.Printf("result: %v\n", result)
}

// create record with selected fields
func createRecordWithSelectedFields() {
	db := connectMySQL()

	type User struct {
		Name      string
		Age       int
		CreatedAt time.Time
	}

	// Create a record and assign a value to the fields specified.
	user := User{Name: "jinzhu", Age: 18, CreatedAt: time.Now()}
	db.Select("Name", "Age", "CreatedAt").Create(&user)
	// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")

	// Create a record and ignore the values for fields passed to omit.
	db.Omit("Name", "Age", "CreatedAt").Create(&user)
	// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")
}

func batchInsert() {
	db := connectMySQL()
	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&users)

	// for _, user := range users {
	// 	user.ID // 1,2,3
	// }

	// batch insert 1000 records
	// var users = []User{{Name: "jinzhu_1"}, ...., {Name: "jinzhu_10000"}}
	// db.CreateInBatches(users, 100) // batch size 100
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New()

	if u.Role == "admin" {
		return errors.New("invalid role")
	}
	return
}

func createHooks() {
	// 实现 gorm 要求的 Hooks 即可, 例如上方的 BeforeCreate, 只要实现了这个方法, 在 Create 时就会自动调用
}
func skipHooks() {
	db := connectMySQL()
	user := User{Name: "jinzhu", Role: "admin"}
	users := []User{{Name: "jinzhu", Role: "admin"}, {Name: "jinzhu2", Role: "user"}}

	//
	db.Session(&gorm.Session{SkipHooks: true}).Create(&user)

	db.Session(&gorm.Session{SkipHooks: true}).Create(&users)

	db.Session(&gorm.Session{SkipHooks: true}).CreateInBatches(users, 100)
}

func createFromMap() {
	db := connectMySQL()
	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "jinzhu", "Age": 18,
	})

	// batch insert from `[]map[string]interface{}{}`
	db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "jinzhu_1", "Age": 18},
		{"Name": "jinzhu_2", "Age": 20},
	})
}

// Upsert / On Conflict
func upsertOnConflict() {
	db := connectMySQL()
	user := User{Name: "jinzhu", Age: 18, Role: "admin"}
	users := []User{{Name: "jinzhu", Age: 18, Role: "admin"}, {Name: "jinzhu2", Age: 20, Role: "user"}}
	// Do nothing on conflict
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)

	// Update columns to default value on `id` conflict
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"role": "user"}),
	}).Create(&users)
	// MERGE INTO "users" USING *** WHEN NOT MATCHED THEN INSERT *** WHEN MATCHED THEN UPDATE SET ***; SQL Server
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE ***; MySQL

	// Use SQL expression
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"count": gorm.Expr("GREATEST(count, VALUES(count))")}),
	}).Create(&users)
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `count`=GREATEST(count, VALUES(count));

	// Update columns to new value on `id` conflict
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
	}).Create(&users)
	// MERGE INTO "users" USING *** WHEN NOT MATCHED THEN INSERT *** WHEN MATCHED THEN UPDATE SET "name"="excluded"."name"; SQL Server
	// INSERT INTO "users" *** ON CONFLICT ("id") DO UPDATE SET "name"="excluded"."name", "age"="excluded"."age"; PostgreSQL
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `name`=VALUES(name),`age`=VALUES(age); MySQL

	// Update all columns to new value on conflict except primary keys and those columns having default values from sql func
	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&users)
	// INSERT INTO "users" *** ON CONFLICT ("id") DO UPDATE SET "name"="excluded"."name", "age"="excluded"."age", ...;
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `name`=VALUES(name),`age`=VALUES(age), ...; MySQL
}
