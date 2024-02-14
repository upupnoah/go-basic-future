package gorm

// https://gorm.io/docs/update.html

func saveAllFields() {
	db := connectMySQL()
	user := User{}
	db.First(&user)

	user.Name = "jinzhu 2"
	user.Age = 100
	user.ID = 111
	db.Save(&user)
	// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;
}

func updateSingleColumn() {
	db := connectMySQL()
	user := User{}

	// Update with conditions
	db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

	// User's ID is `111`:
	db.Model(&user).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

	// Update with conditions and model value
	db.Model(&user).Where("active = ?", true).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;
}

func updateMultipleColumns() {
	db := connectMySQL()
	user := User{}
	// Update attributes with `struct`, will only update non-zero fields
	db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
	// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

	// Update attributes with `map`
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
}

func updateSelectedFields() {
	db := connectMySQL()
	user := User{}
	// Select with Map
	// User's ID is `111`:
	db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET name='hello' WHERE id=111;

	// 忽略某些字段
	db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

	// Select with Struct (select zero value fields)
	db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})
	// UPDATE users SET name='new_name', age=0 WHERE id=111;

	// Select all fields (select all fields include zero value fields)
	db.Model(&user).Select("*").Updates(User{Name: "jinzhu", Role: "admin", Age: 0})

	// Select all fields but omit Role (select all fields include zero value fields)
	db.Model(&user).Select("*").Omit("Role").Updates(User{Name: "jinzhu", Role: "admin", Age: 0})

}

// some hooks
