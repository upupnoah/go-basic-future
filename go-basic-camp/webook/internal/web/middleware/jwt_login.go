package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/web"
)

type LoginJWTMiddleware interface {
	CheckLoginJWT() gin.HandlerFunc
	SetIgnorePaths(paths ...string) LoginJWTMiddleware
}

type LoginJWTMiddlewareBuilder struct {
	ignorePaths map[string]struct{}
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{
		ignorePaths: map[string]struct{}{},
	}
}

func (ljmb *LoginJWTMiddlewareBuilder) CheckLoginJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := ljmb.ignorePaths[ctx.Request.URL.Path]; ok {
			ctx.Next()
			return
		}
		// Authorization 头部, 得到的格式 Bearer token
		authCode := ctx.GetHeader("Authorization")
		if authCode == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		authSegments := strings.SplitN(authCode, " ", 2)
		if len(authSegments) != 2 {
			// 格式不对
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 解析 token
		tokenStr := authSegments[1]
		uc := web.UserClaims{}
		token, err := jwt.ParseWithClaims(
			tokenStr,
			&uc,
			func(token *jwt.Token) (interface{}, error) {
				return web.JWTKey, nil
			})

		if err != nil || !token.Valid {
			// 不正确的 token
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 通过 Claims 检查 UserAgent
		// 对于攻击者来说, 花费的代价更大
		if uc.UserAgent != ctx.Request.UserAgent() {
			// 不正确的 UserAgent, 安全问题
			// 需要加监控
			log.Println("UserAgent not match:", uc.UserAgent, ctx.Request.UserAgent())
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 检查 token 是否过期
		expireTime, err := uc.GetExpirationTime()
		if err != nil {
			// 拿不到过期时间
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if expireTime.Before(time.Now()) {
			// 已经过期
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 检查token是否即将在10分钟内过期
		if time.Until(expireTime.Time) < time.Minute*10 {
			// 更新token的过期时间
			uc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 20)) // 续约 20 分钟
			newToken, err := token.SignedString(web.JWTKey)
			if err != nil {
				// 如果token刷新失败，记录日志
				log.Println("Token refresh failed:", err)
			} else {
				// 如果成功刷新token，则更新响应头
				ctx.Header("x-jwt-token", newToken)
			}
		}
		// 将用户信息添加到上下文中，避免后续操作需要重新解析token
		ctx.Set("user", uc)
	}
}

func (ljmb *LoginJWTMiddlewareBuilder) SetIgnorePaths(paths ...string) LoginJWTMiddleware {
	for _, path := range paths {
		ljmb.ignorePaths[path] = struct{}{}
	}
	return ljmb
}
