package middleware

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginMiddleware interface {
	CheckLogin() gin.HandlerFunc
	SetIgnorePaths(paths ...string) LoginMiddleware
}

type LoginMiddlewareBuilder struct {
	ignorePaths map[string]struct{}
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{
		ignorePaths: map[string]struct{}{
			// "/users/login":  {},
			// "/users/signup": {},
		},
	}
}

// CheckLogin 检查用户是否登录(除了 login 和 signup 之外的路径)
func (lmb *LoginMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := lmb.ignorePaths[ctx.Request.URL.Path]; ok {
			ctx.Next()
			return
		}
		sess := sessions.Default(ctx)
		// 验证一下就可以
		if sess.Get("userId") == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//ctx.Next()
		const timeKey = "update_time"
		val := sess.Get(timeKey)
		updateTime, ok := val.(time.Time)
		// 处于演示效果，整个 session 的过期时间是 1 分钟，所以这里十秒钟刷新一次
		// val == nil 是说明刚登录成功
		// 我们不在登录里面初始化这个 update_time，是因为它属于"刷新"机制，而不属于登录机制
		if val == nil || (ok && time.Since(updateTime) > time.Second*10) {
			sess.Options(sessions.Options{
				MaxAge: 60,
			})
			sess.Set(timeKey, time.Now())
			if err := sess.Save(); err != nil {
				panic(err)
			}
		}
	}
}

func (lmb *LoginMiddlewareBuilder) SetIgnorePaths(paths ...string) LoginMiddleware {
	for _, path := range paths {
		lmb.ignorePaths[path] = struct{}{}
	}
	return lmb
}
