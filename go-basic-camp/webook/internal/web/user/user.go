package user

import (
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewHandler() *Handler {
	const (
		emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		// 和上面比起来，用 ` 看起来就比较清爽
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &Handler{
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

func (h *Handler) RegisterRoutesV1(ug *gin.RouterGroup) {
	//ug := router.Group("/users")
	ug.GET("/profile", h.Profile)
	ug.POST("/login", h.Login)
	ug.POST("/signup", h.Signup)
	ug.POST("/edit", h.Edit)
}

func (h *Handler) Login(c *gin.Context) {
	// todo
}

func (h *Handler) Profile(c *gin.Context) {
	c.String(200, "profile")
}

func (h *Handler) Signup(c *gin.Context) {
	type signUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var req signUpReq
	if err := c.Bind(&req); err != nil {
		return
	}

	matched, err := h.emailExp.MatchString(req.Email)
	if err != nil {
		//c.String(http.StatusInternalServerError, "系统错误!") / / 不要将 err 信息返回给用户
		c.String(http.StatusOK, "系统错误!") // 统一返回200, 表示到达了服务器
		// todo log
		return
	}
	if !matched {
		//c.String(http.StatusBadRequest, "邮箱格式错误!")
		c.String(http.StatusOK, "邮箱格式错误!") // 统一返回200, 表示到达了服务器
		// todo log
		return
	}

	matched, err = h.passwordExp.MatchString(req.Password)
	if err != nil {
		//c.String(http.StatusInternalServerError, "系统错误!") // 不要将 err 信息返回给用户
		c.String(http.StatusOK, "系统错误!") // 统一返回200, 表示到达了服务器
		// todo log
		return
	}
	if !matched {
		//c.String(http.StatusBadRequest, "邮箱格式错误!")
		c.String(http.StatusOK, "密码必须大于8位, 包含数字、特殊字符!") // 统一返回200, 表示到达了服务器
		// todo log
		return
	}

	if req.Password != req.ConfirmPassword {
		fmt.Printf("%+v\n", req)
		c.String(http.StatusOK, "两次密码不一致!") // 统一返回200, 表示到达了服务器
		// todo log
		return
	}

	c.String(http.StatusOK, "注册成功!")
	fmt.Printf("%v\n", req)

	// 下边是数据库操作
}

func (h *Handler) Edit(c *gin.Context) {
	// todo
}
