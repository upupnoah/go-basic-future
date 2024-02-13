package errs

import (
	"fmt"
	"time"
)

// NewErrIndexOutOfRange  创建一个代表下标超出范围的错误
func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("ekit: 下标超出范围, 长度%d, 下标 %d", length, index)
}

// NewErrInvalidType 创建一个代表类型转换失败的错误
func NewErrInvalidType(want string, got any) error {
	return fmt.Errorf("ekit: 类型转换失败，预期类型:%s, 实际值:%#v", want, got)
}

// NewErrInvalidIntervalValue 创建一个代表时间间隔值无效的错误
func NewErrInvalidIntervalValue(interval time.Duration) error {
	return fmt.Errorf("ekit: 无效的间隔时间 %d，预期值应大于 0", interval)
}

// NewErrInvalidMaxIntervalValue 创建一个代表最大重试间隔时间无效的错误
func NewErrInvalidMaxIntervalValue(maxInterval, initialInteval time.Duration) error {
	return fmt.Errorf("ekit: 最大重试间隔时间 [%d]，应大于初始重试的间隔时间 [%d]", maxInterval, initialInteval)
}
