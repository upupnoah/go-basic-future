package slice

// calCapacity 计算切片的容量, 以及是否需要扩容。返回新的容量和是否需要扩容
// 考虑 64， 2048（c/l>=4 -> /2,  c/l >= 2 -> 0.625）
func calCapacity(c, l int) (int, bool) {
	if c <= 64 {
		return c, false
	}
	if c <= 2048 && c/l >= 4 {
		return c / 2, true
	}
	if c > 2048 && c/l >= 2 {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}

	return c, false
}

func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	if n, changed := calCapacity(c, l); !changed {
		return src
	} else {
		s := make([]T, 0, n)
		s = append(s, src...)
		return s
	}
}
