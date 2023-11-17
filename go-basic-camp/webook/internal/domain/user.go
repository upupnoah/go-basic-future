package domain

import "time"

// User 领域对象，是 DDD 中的 entity
type User struct {
	Id       int64
	Email    string
	Password string
	Ctime    time.Time
	Nickname string
	Birthday time.Time
	AboutMe  string
	// ConfirmPassword string // 这个字段不需要
}
