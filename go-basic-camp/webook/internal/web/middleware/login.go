package middleware

import (
	"encoding/gob"
	"net/http"
	"time"

	"github.com/ecodeclub/ekit/set"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// builder 模式方便扩展，以后可以在这里加字段
// 缺点是：不能要求用户的顺序
type LoginMiddlewareBuilder struct {
	publicPaths set.Set[string]
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	s := set.NewMapSet[string](3)
	s.Add("/users/login")
	s.Add("/users/signup")
	return &LoginMiddlewareBuilder{
		publicPaths: s,
	}
}

func (l *LoginMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	gob.Register(time.Time{}) // 注册一个类型，不然会报错：gob: type not registered for interface: time.Time
	// session 中间件
	return func(ctx *gin.Context) {
		// 不需要登录校验的路径
		if l.publicPaths.Exist(ctx.Request.URL.Path) {
			return
		}

		// 从 session 中获取 userId
		sess := sessions.Default(ctx)
		if sess.Get("userId") == nil {
			// 中断，不要往后执行，也就是不要执行后面的业务逻辑
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 固定时间内，刷新 session
		const timeKey = "update_time"
		val := sess.Get(timeKey)
		updateTime, ok := val.(time.Time)
		// none updateTime
		if val == nil || (ok && time.Since(updateTime) > time.Second*10) {
			sess.Options(sessions.Options{
				MaxAge: 60,
			})
			sess.Set(timeKey, time.Now())
			err := sess.Save()
			if err != nil {
				panic(err)
			}
		}
	}
}
