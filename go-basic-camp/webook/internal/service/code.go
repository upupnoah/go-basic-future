package service

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/repository"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/service/sms"
)

var (
	ErrCodeSendTooMany        = repository.ErrCodeSendTooMany
	ErrCodeVerifyTooManyTimes = repository.ErrCodeVerifyTooManyTimes
)

// 短信 Template ID (对应云平台的模板 ID)
const codeTplId = "112234"

type CodeService interface {
	Send(ctx context.Context, biz, phone string) error
	Verify(ctx context.Context, biz, phone, code string) (bool, error)
}

type SMSCodeService struct {
	sms  sms.Service
	repo repository.CodeRepository
}

func NewCodeService(repo repository.CodeRepository, sms sms.Service) CodeService {
	return &SMSCodeService{
		sms:  sms,
		repo: repo,
	}
}

func (svc *SMSCodeService) Send(ctx context.Context, biz, phone string) error {
	code := svc.generateCode()
	fmt.Println()
	// 先 store
	err := svc.repo.Store(ctx, biz, phone, code)
	if err != nil {
		return err
	}
	// store 成功后再发送
	err = svc.sms.Send(ctx, codeTplId, []string{code}, phone)
	if err != nil {
		// 1. 发送失败, 可以删除 store 的数据(用户体验不好)
		// 2. 发送失败, 可以重试 (额外定义一个 retryable 方法)
		return err
	}
	return nil
}

func (svc *SMSCodeService) Verify(ctx context.Context, biz, phone, code string) (bool, error) {
	return svc.repo.Verify(ctx, biz, phone, code)
}

func (svc *SMSCodeService) generateCode() string {
	code := rand.Intn(1000000)
	return fmt.Sprintf("%06d", code)
}
