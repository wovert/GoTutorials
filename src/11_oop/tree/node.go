package tree

import (
	"fmt"
)

type Node struct {
	Value int
	Left, Right *Node
}
// 结构体方法
func (node *Node) Print() {
	fmt.Print(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node.Ignored.")
		return
	}
	node.Value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

// 工厂函数
func CreateNode(value int) *Node {
	// 返回局部变量
	return &Node{Value:value}
}