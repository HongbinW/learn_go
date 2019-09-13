package main

import "fmt"

type Node struct {
	value int
	left,right *Node
}

func (node *Node) print1(){
	fmt.Println(node.value)
}
func (node Node) print2(){
	fmt.Println(node.value)
}

func main() {
	root := Node{value:3}
	root.print1()
	root.print2()

	pRoot := root
	pRoot.print1()
	pRoot.print2()

	nRoot := &root
	nRoot.print1()
	nRoot.print2()


}
