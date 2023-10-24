package slice

import "github.com/upupnoah/go-basic-future/go-basic-camp/ekit/internal/errs"

func Add[T any](src []T, element T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}
	// 先将 src 扩展一个元素
	var zeroValue T
	src = append(src, zeroValue)
	for i := length; i > index; i-- {
		src[i] = src[i-1]
	}
	src[index] = element
	return src, nil
}
