package week1

import "errors"

var ErrIndexOutOfRange = errors.New("index out of range")

// DeleteAt deletes the element at index.
// If the index is out of range, it returns ErrIndexOutOfRange.
func DeleteAt[T any](src []T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index >= length {
		return nil, ErrIndexOutOfRange
	}
	for i := index; i+1 < length; i++ {
		src[i] = src[i+1]
	}
	return src[:length-1], nil
}

// Shrink shrinks the capacity of a slice by calCapacity.
func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	if n, changed := calCapacity(c, l); !changed {
		return src
	} else {
		s := make([]T, l, n)
		s = append(s, src...)
		return s
	}
}

// calCapacity calibrate the new capacity of a slice.
func calCapacity(c, l int) (int, bool) {
	// 容量 <=64 缩不缩都无所谓，因为浪费内存也浪费不了多少
	// 你可以考虑调大这个阈值，或者调小这个阈值
	if c <= 64 {
		return c, false
	}

	// 如果容量大于 2048，但是元素不足一半，
	// 降低为 0.625，也就是 5/8
	// 也就是比一半多一点，和正向扩容的 1.25 倍相呼应
	if c > 2048 && l < c/2 {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}

	// 如果在 2048 以内，并且元素不足 1/4，那么直接缩减为一半
	if c <= 2048 && (c/l > 4) {
		return c / 2, true
	}
	return c, false
}
