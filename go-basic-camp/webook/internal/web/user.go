package web

import (
	"errors"
	"net/http"
	"time"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/domain"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository/cache"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/service"
)

const (
	// 用 ` 看起来就比较清爽, 不然需要转义
	emailRegexPattern = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`

	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`

	phoneRegexPattern = `^1[3-9]\d{9}$`
	//userIdKey            = "user_id"
	bizLogin = "user/login"
)

type UserHandler struct {
	emailRegexExp    *regexp.Regexp
	passwordRegexExp *regexp.Regexp
	phoneRegexExp    *regexp.Regexp
	svc              *service.UserService
	codeService      *service.SMSCodeService
}

// NewUserHandler New UserHandler
func NewUserHandler(svc *service.UserService, codeSvc *service.SMSCodeService) *UserHandler {
	return &UserHandler{
		emailRegexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRegexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		phoneRegexExp:    regexp.MustCompile(phoneRegexPattern, regexp.None),
		svc:              svc,
		codeService:      codeSvc,
	}
}

func (uh *UserHandler) RegisterRoutes(srv *gin.Engine) {
	// srv.POST("/api/user", u.SignUp) // User sign up
	// srv.POST("/api/user/login", u.Login) // User login
	ug := srv.Group("/users")

	ug.GET("/profile", uh.Profile)
	ug.POST("/signup", uh.SignUp)
	ug.POST("/login", uh.LoginJWT)
	ug.POST("/logout", uh.Logout)
	ug.POST("/edit", uh.Edit)
	ug.POST("/login_sms/code/send", uh.SendLoginSMSCode)
	ug.POST("/login_sms", uh.LoginSMS)
}

func (uh *UserHandler) LoginSMS(ctx *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}
	var req Req
	err := ctx.Bind(&req)
	if err != nil {
		return
	}
	ok, err := uh.codeService.Verify(ctx, bizLogin, req.Phone, req.Code)
	if !ok && err == nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 4,
			Msg:  "verification code error!",
		})
		return
	}
	if err == cache.ErrCodeVerifyTooManyTimes {
		ctx.JSON(http.StatusOK, Result{
			Code: 4,
			Msg:  "Verification attempts exceeded!",
		})
		return
	}

	// login or signup
	u, err := uh.svc.FindOrCreate(ctx, req.Phone)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "server error",
		})
		return
	}

	err = uh.setJWTToken(ctx, u.Id)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Msg: "server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, Result{
		Msg: "login success",
	})
}

func (uh *UserHandler) SendLoginSMSCode(ctx *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	// validate phone
	ok, err := uh.phoneRegexExp.MatchString(req.Phone)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "server error",
		})
		return
	}
	if !ok {
		ctx.JSON(http.StatusOK, Result{
			Code: 4,
			Msg:  "Invalid phone number",
		})
		return
	}
	
	err = uh.codeService.Send(ctx, bizLogin, req.Phone)
	if err != nil {
		if err == cache.ErrCodeSendTooMany {
			ctx.JSON(http.StatusOK, Result{
				Code: 5,
				Msg:  "verification code send too many times!",
			})
			return
		}
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, Result{
		Msg: "send success",
	})
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
	if errors.Is(err, service.ErrUserDuplicateUser) {
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
	if err := uh.setJWTToken(ctx, u.Id); err != nil {
		ctx.String(http.StatusOK, "server error")
		return
	}

	ctx.String(http.StatusOK, "login success")
}

func (*UserHandler) setJWTToken(ctx *gin.Context, uid int64) error {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		UserClaims{
			Uid:       uid,
			UserAgent: ctx.GetHeader("User-Agent"),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 20)),
			},
		})
	tokenStr, err := t.SignedString(JWTKey)
	if err != nil {
		return err
	}
	ctx.Header("x-jwt-token", tokenStr)
	return nil
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
	uc, ok := ctx.MustGet("user").(UserClaims)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// u, err := uh.svc.Profile(ctx, uc.Uid)
	u, err := uh.svc.FindById(ctx, uc.Uid)
	if err != nil {
		// 按理说这里的 id 应该是存在的, 如果不存在, 说明有问题
		ctx.String(http.StatusOK, "server error")
	}
	type User struct {
		Nickname string `json:"nickname"`
		Email    string `json:"email"`
		AboutMe  string `json:"aboutMe"`
		Birthday string `json:"birthday"`
	}

	ctx.JSON(http.StatusOK, User{
		Nickname: u.Nickname,
		Email:    u.Email,
		AboutMe:  u.AboutMe,
		Birthday: u.Birthday.Format(time.DateOnly),
	})
}
