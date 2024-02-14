package gorm

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

// https://gorm.io/docs/query.html

// Retrieving a single object
func retrievingSingleObject() {
	db := connectMySQL()
	user := User{}
	// Get the first record ordered by primary key
	db.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// Get one record, no specified order
	db.Take(&user)
	// SELECT * FROM users LIMIT 1;

	// Get last record, ordered by primary key desc
	db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	result := db.First(&user)
	// result.RowsAffected // returns count of records found
	// result.Error        // returns error or nil

	// check error ErrRecordNotFound
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Record not found
		log.Printf("Record not found")
	}

	// --------------------------------------------------------------------

	// works because destination struct is passed in
	db.First(&user)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	// works because model is specified using `db.Model()`
	result1 := map[string]interface{}{}
	db.Model(&User{}).First(&result1)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	// doesn't work
	result2 := map[string]interface{}{}
	db.Table("users").First(&result2)

	// works with Take
	result3 := map[string]interface{}{}
	db.Table("users").Take(&result3)

	// no primary key defined, results will be ordered by first field (i.e., `Code`)
	type Language struct {
		Code string
		Name string
	}
	db.First(&Language{})
	// SELECT * FROM `languages` ORDER BY `languages`.`code` LIMIT 1
}

// 使用主键检索对象
func retrievingObjectWithPrimaryKey() {
	db := connectMySQL()
	user := User{}
	users := []User{}
	db.First(&user, 10) // 传入数字, 默认使用主键检索
	// SELECT * FROM users WHERE id = 10;

	// 使用字符串要小心 SQL 注入
	db.First(&user, "10")
	// SELECT * FROM users WHERE id = 10;

	db.Find(&users, []int{1, 2, 3})
	// SELECT * FROM users WHERE id IN (1,2,3);

	// 使用这种方式避免 SQL 注入
	db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

	var user1 = User{ID: 10}
	db.First(&user1)
	// SELECT * FROM users WHERE id = 10;

	// 使用 Model 灵活性更高（如添加额外的过滤条件、排序等）
	var result User
	db.Model(User{ID: 10}).First(&result)
	// SELECT * FROM users WHERE id = 10;

	// 使用Model方法进行定制化查询
	var res []User
	db.Model(&User{}).Where("age > ?", 18).Order("age desc").Find(&res)

	type User1 struct {
		ID        string         `gorm:"primarykey;size:16"`
		Name      string         `gorm:"size:24"`
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	var user2 = User1{ID: "15"}
	db.First(&user2)
	//  SELECT * FROM `users` WHERE `users`.`id` = '15' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
}
