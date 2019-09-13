package main

import (
	"fmt"
	tree2 "learn_go/struct/tree"
)

type myTreeNode struct {		//组合
	node *tree2.Node //不一定非要是指针
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
	var root tree2.Node
	root = tree2.Node{Value: 3}
	root.Left = &tree2.Node{}
	root.Right = &tree2.Node{5,nil,nil} //因为Left和Right都是指针，所以要赋地址
	root.Right.Left = new(tree2.Node)   //new内建函数，创建一个TreeNode，并返回地址
	root.Left.Left = &tree2.Node{Value: 5,Left:nil,Right:nil}
	root.Left.Right = tree2.CreateTreeNode(2)
	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}


