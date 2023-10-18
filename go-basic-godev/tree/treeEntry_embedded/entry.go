package main

import (
	"fmt"

	"github.com/upupnoah/go-basic-future/go-basic-godev/tree"
)

type myTreeNode struct {
	// node *tree.Node
	*tree.Node // Embedding 内嵌（省略掉名字）
	// 可以直接点出 成员 和 方法
}

// Traverse 可以shadowed method （相当于继承中的重载）
// 但是Go中，不能将子类指针赋值给基类指针（java这样做可以拿到子类重载的方法）
func (myNode *myTreeNode) Traverse() {
	fmt.Println("This is a shadowed method!")
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	// 通过myTreeNode包装一下，就可以postOrder了
	// 但是这样包装的不能自动传入地址， 所以需要一个变量来接收一下
	// myTreeNode{myNode.node.Left}.postOrder()
	// myTreeNode{myNode.node.Right}.postOrder()
	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}
	left.postOrder()
	right.postOrder()
	// myNode.Node.Print()
	myNode.Print()
}

func main() {
	/* root := tree.Node{Value: 3} */
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Right = new(tree.Node)
	root.Right = tree.CreateNode(5)
	fmt.Println("In-order traversal: ")
	root.Traverse()

	fmt.Println("************************")

	fmt.Println("Post-order traversal: ")
	root.postOrder()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println()
	fmt.Println("Node count:", nodeCount)
}
