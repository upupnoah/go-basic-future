package generic

func Sum[T Number](vals ...T) T {
	var res T
	for _, val := range vals {
		res += val
	}
	return res
}

// ~int 表示 int 以及 int 的子类型（衍生类型）

type Number interface {
	~int | int64 | float32 | float64
}

type Integer int
