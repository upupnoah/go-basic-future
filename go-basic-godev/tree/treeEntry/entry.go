package main

import (
	"fmt"
	"github.com/upupnoah/go-basic-future/go-basic/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	// 通过myTreeNode包装一下，就可以postOrder了
	// 但是这样包装的不能自动传入地址， 所以需要一个变量来接收一下
	// myTreeNode{myNode.node.Left}.postOrder()
	// myTreeNode{myNode.node.Right}.postOrder()
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	root := tree.Node{Value: 3}
	// root2 := treeNode{2, nil} //参数少不行
	//root3 := {3, &treeNode{}, &treeNode{}}
	root.Right = new(tree.Node)
	root.Right = tree.CreateNode(999)
	fmt.Println("In-order traversal: ")
	root.Traverse()

	fmt.Println("************************")

	newRoot := myTreeNode{&root}
	fmt.Println("Post-order traversal: ")
	newRoot.postOrder()

}
