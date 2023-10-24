package slice

import "github.com/upupnoah/go-basic-future/go-basic-camp/ekit"

// Max 返回最大值
// 该方法假设你至少会传入一个值
// 在使用 float32 或者 float64 的时候要小心精度问题
func Max[T ekit.RealNumber](ts []T) T {
	res := ts[0]
	for i := 1; i < len(ts); i++ {
		if ts[i] > res {
			res = ts[i]
		}
	}
	return res
}

// Min 返回最小值
// 该方法假设你至少会传入一个值
// 在使用 float32 或者 float64 的时候要小心精度问题
func Min[T ekit.RealNumber](ts []T) T {
	res := ts[0]
	for i := 1; i < len(ts); i++ {
		if ts[i] < res {
			res = ts[i]
		}
	}
	return res
}

// Sum 求和
// 在使用 float32 或者 float64 的时候要小心精度问题
func Sum[T ekit.Number](ts []T) T {
	var res T
	for _, v := range ts {
		res += v
	}
	return res
}
