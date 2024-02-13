package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Using AsciiJSON to Generates ASCII-only JSON with escaped non-ASCII characters.
// 使用 AsciiJSON 生成仅包含转义非 ASCII 字符的 ASCII JSON。

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]any{
			"lang": "GO语言",
			"tag":  "<br>",
			"noah": 111,
			"a":    "a-test",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
