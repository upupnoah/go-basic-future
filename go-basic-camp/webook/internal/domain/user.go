package domain

// User 领域对象，是 DDD 中的 entity
type User struct {
	Email string
	Password string
	// ConfirmPassword string // 这个字段不需要
}

