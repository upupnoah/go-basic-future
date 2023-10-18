package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
// 这里改成comparable即可， 因为comparable是一个接口，所以可以比较两个类型是否相等
type List[T comparable] struct {
	next *List[T]
	val  T
}

func (l *List[T]) insert(val T) {
	t := &List[T]{val: val}
	t.next = l.next
	l.next = t
}

// delete removes the first occurrence of val from l.
func (l *List[T]) delete(val T) {
	for p := l; p != nil; p = p.next {
		if p.next != nil && p.next.val == val {
			p.next = p.next.next
			return
		}
	}
}

// show traverses l and prints its elements to standard output.
func (l List[T]) show() {
	for p := l.next; p != nil; p = p.next {
		fmt.Printf("%v ", p.val)
	}
}

func main() {
	var l List[int]
	l.insert(10)
	l.insert(20)
	l.insert(15)
	l.insert(10)
	l.show()
}
