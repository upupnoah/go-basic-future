package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	key := make([]byte, 32) // 32位字符串长度

	for i := range key {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err) // 处理随机数生成失败
		}
		key[i] = charset[num.Int64()]
	}

	keyString := string(key)
	fmt.Println("密钥:", keyString)
}
