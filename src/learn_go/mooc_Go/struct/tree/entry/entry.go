package main

import (
	"fmt"
	"learn_go/mooc_Go/struct/tree"
)

type myTreeNode struct {		//组合
	node *tree.Node //不一定非要是指针
}

func (myNode *myTreeNode) postOrder(){
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}	//要把这变量单独提出来，不可以直接myTreeNode{myNode.node.Left}.postOrder
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5,nil,nil} //因为Left和Right都是指针，所以要赋地址
	root.Right.Left = new(tree.Node)   //new内建函数，创建一个TreeNode，并返回地址
	root.Left.Left = &tree.Node{Value: 5,Left:nil,Right:nil}
	root.Left.Right = tree.CreateTreeNode(2)
	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	//通过传入函数，可以对每个node，自定义动作（匿名函数）
	nodeCount := 0
	root.TraverseFunc(func(n *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:",nodeCount)
}