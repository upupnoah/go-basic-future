package gorm

type Result struct {
	ID   int
	Name string
	Age  int
}

func rawSql() {
	db := connectMySQL()
	var result Result
	db.Raw("SELECT id, name, age FROM users WHERE id = ?", 3).Scan(&result)

	db.Raw("SELECT id, name, age FROM users WHERE name = ?", "jinzhu").Scan(&result)

	var age int
	db.Raw("SELECT SUM(age) FROM users WHERE role = ?", "admin").Scan(&age)

	var users []User
	db.Raw("UPDATE users SET name = ? WHERE age = ? RETURNING id, name", "jinzhu", 20).Scan(&users)

}
