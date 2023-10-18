package types

import "io"

func UseList() {
	l1 := LinkedList{}
	l1Ptr := &l1
	var l2 LinkedList = *l1Ptr
	println(l2)

	// 这个是 nil
	var l3Ptr *LinkedList
	println(l3Ptr)
}

type List interface {
	Add(index int, val any)
	Append(val any)
	Delete(index int)
}

var _ List = &LinkedList{}

// LinkedList 是一个链表
type LinkedList struct {
	head *node
	//Head *node
}

// Add 添加一个元素
func (l *LinkedList) Add(index int, val any) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Append(val any) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Delete(index int) {
	//TODO implement me
	panic("implement me")
}

type node struct {
	next *node
	// 自引用不用指针会编译错误
	//next node
}

type ListV1[T any] interface {
	Add(index int, val T)
	Append(val T)
	Delete(index int)
}

type LinkedListV1[T any] struct {
	head *nodeV1[T]
}

func (l *LinkedListV1[T]) Add(index int, val T) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedListV1[T]) Append(val T) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedListV1[T]) Delete(index int) {
	//TODO implement me
	panic("implement me")
}

type nodeV1[T any] struct {
	val  T
	next *nodeV1[T]
}

type nodeV2[T io.Closer] struct {
	val  T
	next *nodeV1[T]
}

func (n nodeV2[T]) Use() {
	n.val.Close()
}

func Sum[T Number](vals ...T) T {
	var t T
	for _, ele := range vals {
		t = t + ele
	}
	return t
}

type Integer int

type Number interface {
	~int | uint | uint8
}

func UseTypeP() {
	sum1 := Sum[int](2, 3, 4)
	println(sum1)
	sum2 := Sum[uint](2, 3, 4)
	println(sum2)

	list := &LinkedListV1[string]{}
	list.Append("hello")
}
