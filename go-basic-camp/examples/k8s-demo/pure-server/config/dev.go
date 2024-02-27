//go:build !k8s
package config

// Config dev 配置
var Config = webookConfig{
	DB: DBConfig{
		DSN: "root:root@tcp(localhost:13316)/webook",
	},
	Redis: RedisConfig{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	},
}
