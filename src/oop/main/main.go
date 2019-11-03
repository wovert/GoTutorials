package main

import (
	"fmt"
	"../tree"
)

type myTreeNode struct {
	node *tree.Node
}
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	// var root TreeNode
	// fmt.Println(root) // 0 nil nli

	root := tree.Node {Value: 3}

	// 因为left是指针，所以取地址
	root.Left = &tree.Node {}
	root.Right = &tree.Node {5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	// root.print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	// fmt.Println(root)

	root.Traverse()
	

	// nodes := []TreeNode {
	// 	{value:3},
	// 	{},
	// 	{6, nil, &root},
	// }
	// fmt.Println(nodes)

	fmt.Println()
	my := myTreeNode{&root}
	my.postOrder()
	fmt.Println()
}