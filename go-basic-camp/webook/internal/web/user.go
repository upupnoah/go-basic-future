package web

import (
	"errors"
	"log"
	"net/http"
	"time"

	regexp "github.com/dlclark/regexp2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/domain"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	// 用 ` 看起来就比较清爽, 不然需要转义
	emailRegexPattern = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`

	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	//userIdKey            = "user_id"
	// bizLogin  = "login"
)

type UserHandler struct {
	emailRegexExp    *regexp.Regexp
	passwordRegexExp *regexp.Regexp
	svc              *service.UserService
}

// NewUserHandler New UserHandler
func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		emailRegexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRegexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		svc:              svc,
	}
}

func (uh *UserHandler) RegisterRoutes(srv *gin.Engine) {
	// srv.POST("/api/user", u.SignUp) // User sign up
	// srv.POST("/api/user/login", u.Login) // User login

	srv.POST("/users/signup", uh.SignUp)  // User sign up
	srv.POST("/users/login", uh.LoginJWT) // User login
	srv.POST("/users/edit", uh.Edit)      // User edit
	srv.GET("/users/profile", uh.Profile) // User profile
}

func (uh *UserHandler) SignUp(ctx *gin.Context) {
	type SingUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirm_password"`
		Password        string `json:"password"`
	}
	var req SingUpReq
	// if err, return 400 (ctx.BindJSON)
	if err := ctx.BindJSON(&req); err != nil {
		return
	}

	// validate email
	matched, err := uh.emailRegexExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "System error") // 不暴露具体错误信息
		return
	}
	if !matched {
		ctx.String(http.StatusOK, "Invalid email")
		return
	}
	matched, err = uh.passwordRegexExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "System error")
		return
	}
	if !matched {
		ctx.String(http.StatusOK, "Password is not strong enough. It should contain at least 8 characters")
		return
	}

	// DB
	err = uh.svc.SignUp(ctx.Request.Context(), domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if errors.Is(err, service.ErrUserDuplicateEmail) {
		ctx.String(http.StatusOK, "email duplicate, please use another email")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "server error, signup failed")
	}

	ctx.String(http.StatusOK, "signup success")
}

func (uh *UserHandler) LoginJWT(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.BindJSON(&req); err != nil {
		return
	}
	u, err := uh.svc.Login(ctx.Request.Context(), domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if errors.Is(err, service.ErrInvalidUserOrPassword) {
		ctx.String(http.StatusOK, "invalid email or password, please try again")
		return
	}
	// JWT token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		UserClaims{
			Id:        u.Id,
			UserAgent: ctx.GetHeader("User-Agent"),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 20)),
			},
		})
	tokenStr, err := t.SignedString(JWTKey) // 这个 key 是关键
	if err != nil {
		log.Println("jwt.SignedString error: ", err)
		ctx.String(http.StatusOK, "server error, login failed")
		return
	}
	ctx.Header("x-jwt-token", tokenStr)

	ctx.String(http.StatusOK, "login success")
}

// func (uh *UserHandler) Login(ctx *gin.Context) {
// 	type LoginReq struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}
// 	var req LoginReq
// 	if err := ctx.BindJSON(&req); err != nil {
// 		return
// 	}

// 	// u, err := u.svc.Login(ctx.Request.Context(), domain.User{
// 	// 	Email:    req.Email,
// 	// 	Password: req.Password,
// 	// })
// 	u, err := uh.svc.Login(ctx.Request.Context(), domain.User{
// 		Email:    req.Email,
// 		Password: req.Password,
// 	})
// 	if err == service.ErrInvalidUserOrPassword {
// 		ctx.String(http.StatusOK, "invalid email or password, please try again")
// 		return
// 	}
// 	if err != nil {
// 		ctx.String(http.StatusOK, "server error, login failed")
// 		return
// 	}
// 	session := sessions.Default(ctx)
// 	session.Set("user_id", u.Id)
// 	session.Options(sessions.Options{
// 		MaxAge: 30 * 60, // 30min
// 	})
// 	session.Save()
// 	ctx.String(http.StatusOK, "login success")
// }

func (uh *UserHandler) Logout(ctx *gin.Context) {

}

func (uh *UserHandler) Edit(ctx *gin.Context) {
}

func (uh *UserHandler) Profile(ctx *gin.Context) {
	type ProfileResp struct {
		Email string `json:"email"`
	}
	v, _ := ctx.Get("user")
	userClaims, ok := v.(UserClaims)
	if !ok {
		ctx.String(http.StatusOK, "server error")
		return
	}
	u, err := uh.svc.Profile(ctx, userClaims.Id)
	if err != nil {
		// 按理说这里的 id 应该是存在的, 如果不存在, 说明有问题
		ctx.String(http.StatusOK, "server error")
	}
	ctx.JSON(http.StatusOK, ProfileResp{
		Email: u.Email,
	})
}
