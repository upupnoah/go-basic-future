package web

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"net/http"
	"time"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/domain"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/service"
)

const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	userIdKey            = "userId"
)

type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {

	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		svc:         svc,
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

// func (u *UserHandler) RegisterRoutesV1(ug *gin.RouterGroup) {
// 	//ug := router.Group("/users")
// 	ug.GET("/profile", u.Profile)
// 	// ug.POST("/login", u.Login)
// 	ug.POST("/login", u.LoginJWT) // 使用 JWT
// 	ug.POST("/signup", u.Signup)
// 	ug.POST("/edit", u.Edit)
// }

func (u *UserHandler) RegisterRoutes(engine *gin.Engine) {
	ug := engine.Group("/users")
	ug.GET("/profile", u.Profile)
	// ug.POST("/login", u.Login)
	ug.POST("/login", u.LoginJWT) // 使用 JWT
	ug.POST("/signup", u.Signup)
	ug.POST("/edit", u.Edit)
}

func (u *UserHandler) LoginJWT(ctx *gin.Context) {
	type loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req loginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	user, err := u.svc.Login(ctx, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		ctx.String(http.StatusOK, "账号/邮箱 或 密码错误, 请重试")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	// // 使用 JWT 登录态
	// // 生成 JWT token
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)), // 过期时间暂时设置为 1 分钟
		},
		UserId: user.Id,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims) // 在编码中带上我的数据
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	tokenStr, err := t.SignedString(privateKey)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	ctx.Header("x-jwt-token", tokenStr)
	ctx.String(http.StatusOK, "登录成功")
}

// func (u *UserHandler) Login(ctx *gin.Context) {
// 	type loginReq struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}
// 	var req loginReq
// 	if err := ctx.Bind(&req); err != nil {
// 		return
// 	}
// 	user, err := u.svc.Login(ctx, req.Email, req.Password)
// 	if err == service.ErrInvalidUserOrPassword {
// 		ctx.String(http.StatusOK, "账号/邮箱 或 密码错误, 请重试")
// 		return
// 	}
// 	sess := sessions.Default(ctx)
// 	sess.Set(userIdKey, user.Id)
// 	sess.Options(sessions.Options{
// 		// 60 秒过期
// 		MaxAge: 60,
// 	})
// 	err = sess.Save()
// 	if err != nil {
// 		ctx.String(http.StatusOK, "服务器异常")
// 		return
// 	}
// 	ctx.String(http.StatusOK, "登录成功")
// }

func (u *UserHandler) Profile(ctx *gin.Context) {
	ctx.String(http.StatusOK, "this is your profile")
}

func (u *UserHandler) ProfileJWT(ctx *gin.Context) {
	c, _ := ctx.Get("claims")
	// if !ok {
	// 	// 可以考虑监控住这里，如果出现这种情况，就是代码有问题（因为在中间件 jwt 这一层传入了 claims
	// 	ctx.String(http.StatusOK, "系统错误")
	// 	return
	// }
	claims, ok := c.(*UserClaims)
	if !ok {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	println(claims.UserId)
	// 这边补充 profile 的其他代码
}

func (u *UserHandler) Signup(ctx *gin.Context) {
	type signUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var req signUpReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	matched, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		//c.String(http.StatusInternalServerError, "系统错误!") / / 不要将 err 信息返回给用户
		ctx.String(http.StatusOK, "系统错误!") // 统一返回200, 表示到达了服务器
		// todo log
		return
	}
	if !matched {
		//c.String(http.StatusBadRequest, "邮箱格式错误!")
		ctx.String(http.StatusOK, "邮箱格式错误!") // 统一返回200, 表示到达了服务器
		// todo log
		return
	}
	matched, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		//c.String(http.StatusInternalServerError, "系统错误!") // 不要将 err 信息返回给用户
		ctx.String(http.StatusOK, "系统错误!") // 统一返回200, 表示到达了服务器
		// todo log
		return
	}
	if !matched {
		//c.String(http.StatusBadRequest, "邮箱格式错误!")
		ctx.String(http.StatusOK, "密码必须大于8位, 包含数字、特殊字符!") // 统一返回200, 表示到达了服务器
		// todo log
		return
	}

	if req.Password != req.ConfirmPassword {
		fmt.Printf("%+v\n", req)
		ctx.String(http.StatusOK, "两次密码不一致!") // 统一返回200, 表示到达了服务器
		// todo log
		return
	}

	// 数据库操作
	err = u.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	// 最佳实践
	// if errors.Is(err, service.ErrUserDuplicateEmail) {
	// 	ctx.String(http.StatusOK, "邮箱已经被注册!")
	// 	return
	// }
	if err == service.ErrUserDuplicateEmail {
		ctx.String(http.StatusOK, "邮箱已经被注册!")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	ctx.String(http.StatusOK, "注册成功!")
	// fmt.Printf("%v\n", req)
}

func (u *UserHandler) Edit(ctx *gin.Context) {
	// type Req struct {
	// 	Nickname string `json:"nickname"`
	// 	Birthday string `josn:"birthday"`
	// 	AboutMe  string `json:"aboutMe"`
	// 	// 除了这些字段， 其他字段都不能直接修改
	// 	// 例如：密码、邮箱、手机号
	// }

	// var req Req
	// if err := ctx.Bind(&req); err != nil {
	// 	return
	// }

	// // 校验
	// // 1. nickname 不能超过 20 个字符 && 不能为空
	// // 2. 校验规则取决于产品经理
	// if req.Nickname == "" || len(req.Nickname) > 20 {
	// 	ctx.String(http.StatusOK, "昵称不能为空, 且长度不能超过 20 个字符")
	// 	return
	// }
	// if len(req.AboutMe) > 1024 {
	// 	ctx.String(http.StatusOK, "AboutMe 长度不能超过 1024 个字符")
	// 	return
	// }
	// birthday, err := time.Parse(time.DateOnly, req.Birthday)
	// if err != nil {
	// 	ctx.String(http.StatusOK, "日期格式错误")
	// 	return
	// }

	// u.svc.UpdateProfile(ctx, domain.User{
	// 	Nickname: req.Nickname,
	// 	Birthday: birthday,
	// 	AboutMe:  req.AboutMe,
	// })
	type Req struct {
		Nickname string `json:"nickname"`
		Birthday string `json:"birthday"`
		AboutMe  string `json:"aboutMe"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	// todo after JWT
}

type UserClaims struct {
	jwt.RegisteredClaims
	// 声明自己要放入 token 的数据
	UserId int64
}
