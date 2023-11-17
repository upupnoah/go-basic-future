package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ecodeclub/ekit/set"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/web"
)

type LoginJWTMiddlewareBuilder struct {
	publicPaths set.Set[string]
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {
	s := set.NewMapSet[string](3)
	s.Add("/users/login")
	s.Add("/users/signup")
	return &LoginJWTMiddlewareBuilder{
		publicPaths: s,
	}
}

func (l *LoginJWTMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	// gob.Register(time.Time{}) // 注册一个类型，不然会报错：gob: type not registered for interface: time.Time
	return func(ctx *gin.Context) {
		if l.publicPaths.Exist(ctx.Request.URL.Path) {
			return
		}
		// 使用 JWT 校验
		// 从 header 中获取 token

		tokenHeader := ctx.GetHeader("Authorization")
		if tokenHeader == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 解析 token
		segs := strings.Split(tokenHeader, " ")
		if len(segs) != 2 || segs[0] != "Bearer" {
			// token 格式不对
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenStr := segs[1]
		claims := &web.UserClaims{} // 解析出来的值放到这里
		// 使用 ParseWithClaims 校验 token
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("oiX8LrevZyTBMNQcXS49djcmHVMmxggm"), nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token == nil || !token.Valid || claims.UserId == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// todo 目前这里还不能输出，因为没有校验成功（上面的公钥是错误的）
		fmt.Println("token is valid")
		// 如果校验失败，说明有可能是 token 过期了，那就刷新 token
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))

		// 也可以每 10s 刷新一次
		if time.Until(claims.ExpiresAt.Time) > time.Second*50 {
			s, err := token.SignedString([]byte("oiX8LrevZyTBMNQcXS49djcmHVMmxggm"))
			if err != nil {
				// 记录日志，续约失败，但是不要中断
				log.Println("jwt 续约失败")
			}
			ctx.Header("x-jwt-token", s) // 返回前端
		}

		ctx.Set("claims", claims) // 将 claims 放到 context 中

		// jwt 和 session 混用是主流
		// jwt 登录校验
		// session 存敏感信息
	}
}
