package generic

// generic for interface

type MyList[T any] interface {
	Add(index int, val T)
	Append(val T)
}

// generic for struct

type LinkedList[T any] struct {
	head *node[T]
}

type node[T any] struct {
	val T
}
