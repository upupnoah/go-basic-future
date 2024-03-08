package tencent

import (
	"context"
	"fmt"

	"github.com/ecodeclub/ekit"
	"github.com/ecodeclub/ekit/slice"
	tencentSMS "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms

	mysms "github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/service/sms"
)

type Service struct {
	client   *tencentSMS.Client
	appId    *string
	signName *string
}

func (s *Service) Send(ctx context.Context, templateId string, templateArgs []string, numbers ...string) error {
	request := tencentSMS.NewSendSmsRequest()

	/* 短信应用ID: 短信SdkAppId在 [短信控制台] 添加应用后生成的实际SdkAppId，示例如1400006666 */
	// 应用 ID 可前往 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage) 查看
	request.SmsSdkAppId = s.appId

	/* 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名 */
	// 签名信息可前往 [国内短信](https://console.cloud.tencent.com/smsv2/csms-sign) 或 [国际/港澳台短信](https://console.cloud.tencent.com/smsv2/isms-sign) 的签名管理查看
	request.SignName = s.signName

	// 模板 ID 可前往 [国内短信](https://console.cloud.tencent.com/smsv2/csms-template) 或 [国际/港澳台短信](https://console.cloud.tencent.com/smsv2/isms-template) 的正文模板管理查看
	request.TemplateId = ekit.ToPtr(templateId)
	// 模板参数: 模板参数的个数需要与 TemplateId 对应模板的变量个数保持一致，若无模板参数，则设置为空
	request.TemplateParamSet = toStringPtrSlice(templateArgs)
	// 下发手机号码，采用 E.164 标准，+[国家或地区码][手机号]
	request.PhoneNumberSet = toStringPtrSlice(numbers)

	request.SetContext(ctx) // ctx 继续往下传

	response, err := s.client.SendSms(request)
	// if _, ok := err.(*errors.TencentCloudSDKError); ok {
	// 	log.Printf("An API error has returned: %s", err)
	// 	return err
	// }
	if err != nil {
		return err
	}
	for _, status := range response.Response.SendStatusSet {
		if status.Code == nil || *status.Code != "Ok" {
			return fmt.Errorf("send sms failed, code: %s, reason: %s",
				*status.Code, *status.Message)
		}
	}
	return nil
}

func (s *Service) SendV1(ctx context.Context, templateId string, templateArgs any, numbers ...string) error {
	request := tencentSMS.NewSendSmsRequest()

	/* 短信应用ID: 短信SdkAppId在 [短信控制台] 添加应用后生成的实际SdkAppId，示例如1400006666 */
	// 应用 ID 可前往 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage) 查看
	request.SmsSdkAppId = s.appId

	/* 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名 */
	// 签名信息可前往 [国内短信](https://console.cloud.tencent.com/smsv2/csms-sign) 或 [国际/港澳台短信](https://console.cloud.tencent.com/smsv2/isms-sign) 的签名管理查看
	request.SignName = s.signName

	// 模板 ID 可前往 [国内短信](https://console.cloud.tencent.com/smsv2/csms-template) 或 [国际/港澳台短信](https://console.cloud.tencent.com/smsv2/isms-template) 的正文模板管理查看
	request.TemplateId = ekit.ToPtr(templateId)
	// 模板参数: 模板参数的个数需要与 TemplateId 对应模板的变量个数保持一致，若无模板参数，则设置为空
	request.TemplateParamSet = toStringPtrSlice(templateArgs.([]string))
	// 下发手机号码，采用 E.164 标准，+[国家或地区码][手机号]
	request.PhoneNumberSet = toStringPtrSlice(numbers)

	request.SetContext(ctx) // ctx 继续往下传

	response, err := s.client.SendSms(request)
	// if _, ok := err.(*errors.TencentCloudSDKError); ok {
	// 	log.Printf("An API error has returned: %s", err)
	// 	return err
	// }
	if err != nil {
		return err
	}
	for _, status := range response.Response.SendStatusSet {
		if status.Code == nil || *status.Code != "Ok" {
			return fmt.Errorf("send sms failed, code: %s, reason: %s",
				*status.Code, *status.Message)
		}
	}
	return nil
}

func (s *Service) SendV2(ctx context.Context, templateId string, templateArgs []mysms.NamedArg, numbers ...string) error {
	request := tencentSMS.NewSendSmsRequest()

	/* 短信应用ID: 短信SdkAppId在 [短信控制台] 添加应用后生成的实际SdkAppId，示例如1400006666 */
	// 应用 ID 可前往 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage) 查看
	request.SmsSdkAppId = s.appId

	/* 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名 */
	// 签名信息可前往 [国内短信](https://console.cloud.tencent.com/smsv2/csms-sign) 或 [国际/港澳台短信](https://console.cloud.tencent.com/smsv2/isms-sign) 的签名管理查看
	request.SignName = s.signName

	// 模板 ID 可前往 [国内短信](https://console.cloud.tencent.com/smsv2/csms-template) 或 [国际/港澳台短信](https://console.cloud.tencent.com/smsv2/isms-template) 的正文模板管理查看
	request.TemplateId = ekit.ToPtr(templateId)
	// 模板参数: 模板参数的个数需要与 TemplateId 对应模板的变量个数保持一致，若无模板参数，则设置为空

	// []NamedArg -> []*string
	request.TemplateParamSet = slice.Map[mysms.NamedArg, *string](templateArgs, func(idx int, src mysms.NamedArg) *string {
		return &src.Val
	})
	
	// 下发手机号码，采用 E.164 标准，+[国家或地区码][手机号]
	request.PhoneNumberSet = toStringPtrSlice(numbers)

	request.SetContext(ctx) // ctx 继续往下传

	response, err := s.client.SendSms(request)
	// if _, ok := err.(*errors.TencentCloudSDKError); ok {
	// 	log.Printf("An API error has returned: %s", err)
	// 	return err
	// }
	if err != nil {
		return err
	}
	for _, status := range response.Response.SendStatusSet {
		if status.Code == nil || *status.Code != "Ok" {
			return fmt.Errorf("send sms failed, code: %s, reason: %s",
				*status.Code, *status.Message)
		}
	}
	return nil
}

func toStringPtrSlice(strSlice []string) []*string {
	return slice.Map[string, *string](strSlice, func(idx int, src string) *string {
		return &src
	})
}

func NewService(client *tencentSMS.Client, appId string, signName string) *Service {
	return &Service{
		client:   client,
		appId:    ekit.ToPtr(appId),
		signName: ekit.ToPtr(signName),
	}
}
