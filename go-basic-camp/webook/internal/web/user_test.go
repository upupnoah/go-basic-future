package web

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/domain"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/service"
	svcmocks "github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/service/mocks"
	"go.uber.org/mock/gomock"
)

func TestUserHandlerV1_SignUp(t *testing.T) {
	const signupURL = "/users/signup"
	testCases := []struct {
		name string
		// 准备 mock, 用到了 userService 和 codeService
		mock func(ctrl *gomock.Controller) (service.UserService, service.CodeService)

		// 输入, 因为 request 的构造过程可能很复杂
		// 因此定义一个函数, 用于构造 request
		reqBuilder func(t *testing.T) *http.Request
		wantCode   int
		wantBody   string
	}{
		{
			name: "test1: signup success",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				userService := svcmocks.NewMockUserService(ctrl)
				userService.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(nil)
				// 在 signup 中, 没有用到 codeService, 因此不需要模拟调用
				codeService := svcmocks.NewMockCodeService(ctrl)
				return userService, codeService
			},
			reqBuilder: func(t *testing.T) *http.Request {
				body := bytes.NewBuffer([]byte(`{"email":"test@gmail.com","password":"test0x3f!!","confirm_password":"test0x3f!!"}`))
				req, err := http.NewRequest(http.MethodPost, signupURL, body)
				req.Header.Set("Content-Type", "application/json")
				if err != nil {
					t.Fatal(err)
				}
				return req
			},
			wantCode: http.StatusOK,
			wantBody: "signup success",
		},
		{
			name: "test2: not json",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				// 因为根本没有跑到 signup 那里，所以直接返回 nil 都可以
				return nil, nil
			},
			reqBuilder: func(t *testing.T) *http.Request {
				// 准备一个错误的JSON 串
				body := bytes.NewBuffer([]byte(`{"email":"test@gmail.com",`))
				req, err := http.NewRequest(http.MethodPost, signupURL, body)
				req.Header.Set("Content-Type", "application/json")
				if err != nil {
					t.Fatal(err)
				}
				return req
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "test3: email format error",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				// 因为根本没有跑到 signup 那里，所以直接返回 nil 都可以
				return nil, nil
			},
			reqBuilder: func(t *testing.T) *http.Request {
				// 准备一个不合法的邮箱
				body := bytes.NewBuffer([]byte(`{"email":"test@","password":"test0x3f!!","confirm_password":"test0x3f!!"}`))
				req, err := http.NewRequest(http.MethodPost, signupURL, body)
				req.Header.Set("Content-Type", "application/json")
				if err != nil {
					t.Fatal(err)
				}
				return req
			},
			wantCode: http.StatusOK,
			wantBody: "Invalid email",
		},
		{
			name: "test4: password and confirm password are not the same",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				// 因为根本没有跑到 signup 那里，所以直接返回 nil 都可以
				return nil, nil
			},
			reqBuilder: func(t *testing.T) *http.Request {
				// 准备一个不合法的邮箱
				body := bytes.NewBuffer([]byte(`{"email":"test@gmail.com","password":"!!test0x3f","confirm_password":"test0x3f!!"}`))
				req, err := http.NewRequest(http.MethodPost, signupURL, body)
				req.Header.Set("Content-Type", "application/json")
				if err != nil {
					t.Fatal(err)
				}
				return req
			},
			wantCode: http.StatusOK,
			wantBody: "Password and confirm password are not the same",
		},
		{
			name: "test5: password format error",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				// 因为根本没有跑到 signup 那里，所以直接返回 nil 都可以
				return nil, nil
			},
			reqBuilder: func(t *testing.T) *http.Request {
				// 准备一个不合法的邮箱
				body := bytes.NewBuffer([]byte(`{"email":"test@gmail.com","password":"hello","confirm_password":"hello"}`))
				req, err := http.NewRequest(http.MethodPost, signupURL, body)
				req.Header.Set("Content-Type", "application/json")
				if err != nil {
					t.Fatal(err)
				}
				return req
			},
			wantCode: http.StatusOK,
			wantBody: "Password is not strong enough. It should contain at least 8 characters",
		},
		{
			name: "test6: email duplicate",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				userService := svcmocks.NewMockUserService(ctrl)
				userService.EXPECT().SignUp(gomock.Any(), gomock.Any()).
					// 模拟返回邮箱冲突的异常
					Return(service.ErrUserDuplicateUser)

				// 在 signup 这个接口里面，并没有用到的 codeService，
				// 所以什么不需要准备模拟调用
				codeService := svcmocks.NewMockCodeService(ctrl)
				return userService, codeService
			},
			reqBuilder: func(t *testing.T) *http.Request {
				body := bytes.NewBuffer([]byte(`{"email":"test@gmail.com","password":"test0x3f!!","confirm_password":"test0x3f!!"}`))
				req, err := http.NewRequest(http.MethodPost, signupURL, body)
				req.Header.Set("Content-Type", "application/json")
				if err != nil {
					t.Fatal(err)
				}
				return req
			},
			wantCode: http.StatusOK,
			wantBody: "email duplicate, please use another email",
		},
		{
			name: "test7: server error",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				userService := svcmocks.NewMockUserService(ctrl)
				userService.EXPECT().SignUp(gomock.Any(), gomock.Any()).
					// 注册失败，系统本身的异常
					Return(errors.New("模拟系统异常"))

				// 在 signup 这个接口里面，并没有用到的 codeService，
				// 所以什么不需要准备模拟调用
				codeService := svcmocks.NewMockCodeService(ctrl)
				return userService, codeService
			},
			reqBuilder: func(t *testing.T) *http.Request {
				body := bytes.NewBuffer([]byte(`{"email":"test@gmail.com","password":"test0x3f!!","confirm_password":"test0x3f!!"}`))
				req, err := http.NewRequest(http.MethodPost, signupURL, body)
				req.Header.Set("Content-Type", "application/json")
				if err != nil {
					t.Log(err)
				}
				return req
			},
			wantCode: http.StatusOK,
			wantBody: "System error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userService, codeService := tc.mock(ctrl)

			handler := NewUserHandler(userService, codeService)
			server := gin.Default()
			handler.RegisterRoutes(server)

			// 构造 request
			req := tc.reqBuilder(t)
			// 创建一个 recorder, 用于记录服务端的响应
			recorder := httptest.NewRecorder()
			// 服务端处理请求, 并将响应写入 recorder
			server.ServeHTTP(recorder, req)

			assert.Equal(t, tc.wantCode, recorder.Code)
			assert.Equal(t, tc.wantBody, recorder.Body.String())
		})
	}
}

// mock 的流程
// 1. 创建一个控制器
// 2. 创建一个 mock 对象
// 3. 设置预期, 传入的参数, 返回的结果
// 4. 调用
// 5. 验证预期

func TestMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userService := svcmocks.NewMockUserService(ctrl)
	userService.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(errors.New("error!!!"))

	// 调用
	err := userService.SignUp(context.TODO(), domain.User{}).Error()
	if err == "error!!!" {
		t.Log("error!!!")
	}
}
