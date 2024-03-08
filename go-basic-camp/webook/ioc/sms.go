package ioc

import (
	"os"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tencentSMS "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/service/sms"
	localsms "github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/service/sms/local_sms"
	"github.com/upupnoah/go-basic-future/go-basic-camp/webook/internal/service/sms/tencent"
)

func InitSMSService() sms.Service {
	// 基于本地的短信服务(测试用)
	return localsms.NewService()

	// 可以切换成腾讯/阿里云的短信服务
	// return initTencentSMSService()
	// return initAliSMSService()
}

func initTencentSMSService() sms.Service {
	secretId, ok := os.LookupEnv("SMS_SECRET_ID")
	if !ok {
		panic("找不到腾讯 SMS 的 secret id")
	}
	secretKey, ok := os.LookupEnv("SMS_SECRET_KEY")
	if !ok {
		panic("找不到腾讯 SMS 的 secret key")
	}
	c, err := tencentSMS.NewClient(
		common.NewCredential(secretId, secretKey),
		"ap-nanjing",
		profile.NewClientProfile(),
	)
	if err != nil {
		panic(err)
	}
	return tencent.NewService(c, "10086", "Noah-Future")
}

// func InitAliSMSService() sms.Service {

// }