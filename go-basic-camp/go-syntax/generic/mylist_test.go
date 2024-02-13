package generic

import (
	"testing"
)

type myList[T any] struct {
	data []T
}

func (m myList[T]) Add(index int, val T) {
}

func (m myList[T]) Append(val T) {
	m.data = append(m.data, val)
}

func TestMyList(t *testing.T) {
	var l MyList[int]
	l = myList[int]{}
	l.Append(1)
	//mylist := myList[int]{}
	//mylist.Append(1)
	// mylist.Append(1.3)
	//mylist.Append("hello")
	//fmt.Println(mylist)
}
