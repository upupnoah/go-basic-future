package generic

// T 是一个类型参数
// any 是一个类型约束，表示任意类型

type Pair[T any] struct {
	A, B T
}

func Swap[T any](p Pair[T]) Pair[T] {
	return Pair[T]{p.B, p.A}
}

type Stringer interface {
	String() string
}

type MyType[T Stringer] struct {
	value T
}

type Person struct {
	name string
}

// implement Stringer
func (p Person) String() string {
	return p.name
}
