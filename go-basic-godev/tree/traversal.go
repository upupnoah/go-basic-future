package tree

import "fmt"

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node) // 不管 f 想干嘛我们先不管， 具体看f如何是实现的，想干嘛就干嘛
	node.Right.TraverseFunc(f)
}
