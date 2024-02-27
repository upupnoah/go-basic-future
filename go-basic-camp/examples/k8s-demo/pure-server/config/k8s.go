//go:build k8s
package config

// Config k8s配置
var Config = webookConfig{
	// 具体看 service 的 port, pod 之间相互访问需要通过 service
	// k8s 中的 pod ip 地址在重新部署的时候可能会变
	// 但 service 提供了一个不变的地址 和 DNS 名称
	DB: DBConfig{
		DSN: "root:root@tcp(pure-mysql:11306)/webook",
	},
	Redis: RedisConfig{
		Addr:     "pure-redis:11379",
		Password: "",
		DB:       0,
	},
}
