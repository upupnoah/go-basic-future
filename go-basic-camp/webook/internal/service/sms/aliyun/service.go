package aliyun

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	aliSMS "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/service/sms"
)

type Service struct {
	client   *aliSMS.Client
	signName *string
}

func (s *Service) Send(ctx context.Context, templateId string, templateArgs []string, numbers ...string) error {
	phoneNumbers := strings.Join(numbers, ",")
	templateParam := strings.Join(templateArgs, ",")
	sendSmsRequest := &aliSMS.SendSmsRequest{
		SignName:      s.signName,
		TemplateCode:  tea.String(templateId),
		PhoneNumbers:  tea.String(phoneNumbers),
		TemplateParam: tea.String(templateParam),
	}
	runtime := &util.RuntimeOptions{}
	res, err := s.client.SendSmsWithOptions(sendSmsRequest, runtime)
	if res.Body.Code != nil && *res.Body.Code != "OK" {
		log.Printf("send sms failed, code: %s, reason: %s", *res.Body.Code, *res.Body.Message)
		return err
	}
	return err
}

func (s *Service) SendV1(ctx context.Context, templateId string, templateArgs any, numbers ...string) error {
	bcode, err := json.Marshal(templateArgs.(map[string]string))
	if err != nil {
		return err
	}
	phoneNumbers := strings.Join(numbers, ",")
	sendSmsRequest := &aliSMS.SendSmsRequest{
		SignName:      s.signName,
		TemplateCode:  tea.String(templateId),
		PhoneNumbers:  tea.String(phoneNumbers),
		TemplateParam: tea.String(string(bcode)),
	}
	runtime := &util.RuntimeOptions{}
	res, err := s.client.SendSmsWithOptions(sendSmsRequest, runtime)
	if res.Body.Code != nil && *res.Body.Code != "OK" {
		log.Printf("send sms failed, code: %s, reason: %s", *res.Body.Code, *res.Body.Message)
		return err
	}
	return err
}

func (s *Service) SendV2(ctx context.Context, templateId string, templateArgs []sms.NamedArg, numbers ...string) error {
	phoneNumbers := strings.Join(numbers, ",")
	argsMap := make(map[string]string, len(templateArgs))
	for _, v := range templateArgs {
		argsMap[v.Name] = v.Val
	}
	bcode, err := json.Marshal(argsMap)
	if err != nil {
		return err
	}
	sendSmsRequest := &aliSMS.SendSmsRequest{
		SignName:      s.signName,
		TemplateCode:  tea.String(templateId),
		PhoneNumbers:  tea.String(phoneNumbers),
		TemplateParam: tea.String(string(bcode)),
	}
	runtime := &util.RuntimeOptions{}
	res, err := s.client.SendSmsWithOptions(sendSmsRequest, runtime)
	if res.Body.Code != nil && *res.Body.Code != "OK" {
		log.Printf("send sms failed, code: %s, reason: %s", *res.Body.Code, *res.Body.Message)
		return err
	}
	return err
}

func NewService(client *aliSMS.Client) *Service {
	return &Service{
		client: client,
	}
}
