package localsms

import (
	"context"
	"fmt"
)

type Service struct {
}

// 生成的验证码在 templateArgs 中
// 这里可以认为模版中就一个{}, 传入的验证码就是模版参数
func (s *Service) Send(ctx context.Context, templateId string, templateArgs []string, numbers ...string) error {
	fmt.Println("sms is:", templateArgs)
	return nil
}

func NewService() *Service {
	return &Service{}
}
