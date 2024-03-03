package sms

import "context"

type Service interface {
	// 写成这样是因为加参数的时候, 编译器能报错提示哪里需要修改
	// 也可以定义一个 Request 结构体, 里面定义字段(好处是可以定义默认值)
	// 根据腾讯云 sms 的文档进行抽象
	// appid 和 sign 是固定的, 所以不需要传入, 初始化的时候直接传入即可

	// templateArgs 用 []string:
	// - 传入[]string, 不同的平台(ali, tencent)需要转换成对应的类型
	Send(ctx context.Context, templateId string, templateArgs []string, numbers ...string) error
}

type ServiceV1 interface {
	// templateArgs 用 any:
	// - 传入的参数可以是任意类型, 但是需要在内部进行类型断言
	// - 调用者需要知道传入的参数类型(!!!)
	SendV1(ctx context.Context, templateId string, templateArgs any, numbers ...string) error
}

type ServiceV2 interface {
	// templateArgs 用 []namedArg:
	// - 传入的参数是一个结构体, 里面包含了参数名和参数值
	// - 实现 interface 时方便转成对应的类型
	SendV2(ctx context.Context, templateId string, templateArgs []NamedArg, numbers ...string) error
}

type NamedArg struct {
	Name string
	Val  string
}
