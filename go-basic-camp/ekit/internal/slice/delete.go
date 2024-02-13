package slice

import "github.com/upupnoah/go-basic-future/go-basic-camp/ekit/internal/errs"

// Delete 删除指定下标的元素, 返回删除后的切片和被删除的元素
func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zero T
		//return nil, *new(T), errs.NewErrIndexOutOfRange(length, index)
		return nil, zero, errs.NewErrIndexOutOfRange(length, index)
	}
	res := src[index]
	for i := index; i+1 < length; i++ {
		src[i] = src[i+1]
	}
	// 去掉最后一个重复元素
	src = src[:length-1]
	return src, res, nil
}
