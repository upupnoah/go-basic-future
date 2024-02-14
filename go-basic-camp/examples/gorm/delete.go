package gorm

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// https://gorm.io/docs/delete.html

type Email struct {
	ID   int
	Name string
}

// Delete Hooks
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	if u.Role == "admin" {
		return errors.New("admin user not allowed to delete")
	}
	return
}

func deleteRecord() {
	db := connectMySQL()

	email := Email{ID: 10}

	// Email's ID is `10`
	db.Delete(&email)
	// DELETE from emails where id = 10;

	// Delete with additional conditions
	db.Where("name = ?", "jinzhu").Delete(&email)
	// DELETE from emails where id = 10 AND name = "jinzhu";

}

func batchDelete() {
	db := connectMySQL()
	db.Where("email LIKE ?", "%jinzhu%").Delete(&Email{})
	// DELETE from emails where email LIKE "%jinzhu%";

	db.Delete(&Email{}, "email LIKE ?", "%jinzhu%")
	// DELETE from emails where email LIKE "%jinzhu%";
}

// Return deleted data, only works for database support Returning
func returnDataFromDeleteRows() {
	DB := connectMySQL()
	// return all columns
	var users []User
	DB.Clauses(clause.Returning{}).Where("role = ?", "admin").Delete(&users)
	// DELETE FROM `users` WHERE role = "admin" RETURNING *
	// users => []User{{ID: 1, Name: "jinzhu", Role: "admin", Salary: 100}, {ID: 2, Name: "jinzhu.2", Role: "admin", Salary: 1000}}

	// return specified columns
	DB.Clauses(clause.Returning{Columns: []clause.Column{{Name: "name"}, {Name: "salary"}}}).Where("role = ?", "admin").Delete(&users)
	// DELETE FROM `users` WHERE role = "admin" RETURNING `name`, `salary`
	// users => []User{{ID: 0, Name: "jinzhu", Role: "", Salary: 100}, {ID: 0, Name: "jinzhu.2", Role: "", Salary: 1000}}

}

// Soft Delete
// If your model includes a gorm.DeletedAt field (which is included in gorm.Model), it will get soft delete ability automatically!
// 如果您的模型包含gorm.DeletedAt字段（该字段已包含在gorm.Model中），它将自动获得软删除功能！
func softDelete() {
	db := connectMySQL()
	// user's ID is `111`
	user := User{ID: 111}
	db.Delete(&user)
	// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

	// Batch Delete
	db.Where("age = ?", 20).Delete(&User{})
	// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

	// Soft deleted records will be ignored when querying
	db.Where("age = 20").Find(&user)
	// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;

}

func gormModel() {
	type User3 struct {
		gorm.Model // 嵌入了一些字段
		Name string
	}
	u := User3{Name: "jinzhu"}
	db := connectMySQL()

	db.Create(&u)

	db.Delete(&u)
}
