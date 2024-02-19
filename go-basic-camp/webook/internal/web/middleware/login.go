package middleware

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginMiddlewareBuilder struct {
	ignorePaths map[string]struct{}
}

type LoginMiddleware interface {
	CheckLogin() gin.HandlerFunc
	SetIgnorePaths(paths ...string) LoginMiddleware
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
		session := sessions.Default(ctx)
		if _, ok := lmb.ignorePaths[ctx.Request.URL.Path]; ok {
			ctx.Next()
			return
		}
		id := session.Get("user_id")
		// 未登录, 无权继续访问
		if id == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			// ctx.Redirect(http.StatusFound, "/users/login") // 重定向到登录页面

			return // 到此为止, 不再执行后续的中间件
		}
		session.Set("user_id", id)
		updateTime := session.Get("update_time")
		now := time.Now()
		if updateTime == nil { // 第一次登录
			session.Set("update_time", now)
			session.Options(sessions.Options{
				MaxAge: 30 * 60,
			})
			if err := session.Save(); err != nil {
				panic(err)
			}
			return
		}
		updateTimeVal, ok := updateTime.(time.Time)
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		// 维护 session 的活跃状态, 当前时间与上次更新时间相差超过 1 分钟, 则更新
		if now.Sub(updateTimeVal) > time.Minute {
			session.Set("update_time", now)
			session.Save()
			return
		}
		session.Set("update_time", now)

		// 通行
		ctx.Next()
	}
}

func (lmb *LoginMiddlewareBuilder) SetIgnorePaths(paths ...string) LoginMiddleware {
	for _, path := range paths {
		lmb.ignorePaths[path] = struct{}{}
	}
	return lmb
}
