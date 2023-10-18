package queue

// Queue An FIFO queue.
type Queue []interface{}

// Push 任意类型
//func (q *Queue) Push(v interface{}) {
//	*q = append(*q, v)
//}
//
//func (q *Queue) Pop() interface{} {
//	head := (*q)[0]
//	*q = (*q)[1:]
//	return head
//}

// Push the element into the queue.
// 		e.g.  q.Push(123)
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pop element from head.
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	// head是一个interface{} 类型， 如何限定成int呢？
	return head.(int) //可以将interface中强制转换int的值， 然后拿出来
}

// IsEmpty Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
