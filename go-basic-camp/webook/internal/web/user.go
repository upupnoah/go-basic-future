package web

import (
	"fmt"
	"net/http"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/domain"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/service"
)

type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		// 和上面比起来，用 ` 看起来就比较清爽
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		svc:         svc,
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

func (u *UserHandler) RegisterRoutesV1(ug *gin.RouterGroup) {
	//ug := router.Group("/users")
	ug.GET("/profile", u.Profile)
	ug.POST("/login", u.Login)
	ug.POST("/signup", u.Signup)
	ug.POST("/edit", u.Edit)
}

func (u *UserHandler) RegisterRoutes(engine *gin.Engine) {
	ug := engine.Group("/users")
	ug.GET("/profile", u.Profile)
	ug.POST("/login", u.Login)
	ug.POST("/signup", u.Signup)
	ug.POST("/edit", u.Edit)
}

func (u *UserHandler) Login(ctx *gin.Context) {
	// todo
}

func (u *UserHandler) Profile(ctx *gin.Context) {
	ctx.String(200, "profile")
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
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	ctx.String(http.StatusOK, "注册成功!")
	// fmt.Printf("%v\n", req)
}

func (h *UserHandler) Edit(c *gin.Context) {
	// todo
}
