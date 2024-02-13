package generic

import (
	"cmp"
)

// comparable
// 可比较性

func Max[T cmp.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}
