package tree

import "fmt"

// Node 结构体定义
// 改成大写则表示 public
type Node struct {
	Value       int
	Left, Right *Node
}

// CreateNode 虽然没有构造方法， 但是Go中可以写 工厂函数
func CreateNode(value int) *Node {
	// 在c++中， 返回一个局部变量的地址会出问题
	// 但是在go语言中，可以返回一个局部变量的地址给别人使用
	return &Node{Value: value}
}

// Print 为结构体写方法
// 括号中的是 接收者， 表示 print() 是给 node 来接收的
func (node Node) Print() {
	fmt.Printf("%v ", node.Value)
}

// SetValue go语言中参数是传值的，因此需要加一个*， 传递指针
// go语言的之间引用和间接引用都是. 不需要 ->
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	// nil指针虽然不报错， 但是不能将这个value拿出来... 因此上面要return
	node.Value = value
}
