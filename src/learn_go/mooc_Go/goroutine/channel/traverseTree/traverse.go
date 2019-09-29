package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func (node *TreeNode) TraverseFunc (f func(*TreeNode)){		//函数式编程
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
func (node *TreeNode)traverseWithChannel() chan *TreeNode{	//将node用channel传回去
	out := make(chan *TreeNode)
	go func() {
		node.TraverseFunc(func(node *TreeNode) {
			out <- node
		})
		close(out)
	}()
	return out
}

func main(){
	root := TreeNode{
		Val:   4,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: &TreeNode{
				Val:	6,
			},
		},
	}
	c := root.traverseWithChannel()
	maxNode := 0
	for node := range c{
		fmt.Println(node.Val)
		if node.Val > maxNode{
			maxNode = node.Val
		}
	}
	fmt.Println("Max node value:",maxNode)
}