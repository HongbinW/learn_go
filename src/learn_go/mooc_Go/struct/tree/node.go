package tree

import "fmt"

type Node struct {
	Value int
	Left,Right *Node
}

func CreateTreeNode(Value int) *Node {
	return &Node{Value: Value} //虽然是局部变量，但也可以给别人用。不需要知道该变量分配在堆上还是栈上
}

func (node Node) Print(){ //node是一个接收者，类似print(node treeNode)，传递的都是值。写前面是struct的方法
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(Value int)  {
	if node == nil {
		fmt.Println("Setting nil node. Ignore")
		return
	}
	node.Value = Value;
}
//print和setValue区别：print参数传过来的就是一个值，而setValue传过来的是一个地址。而跟传过来的对象没关系，即使传过来的是个地址，该是值还是值，该是地址还是地址
//值传递修改不会改原来内容

func main() {
	var root Node
	root = Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5,nil,nil} //因为Left和Right都是指针，所以要赋地址
	root.Right.Left = new(Node)   //new内建函数，创建一个TreeNode，并返回地址
	root.Left.Left = &Node{Value: 5,Left:nil,Right:nil}
	root.Left.Right = CreateTreeNode(2)
	root.Traverse()
	//
	//nodes := []Node{
	//	{Value:3},
	//	{},
	//	{6,nil,&root},
	//}
	//fmt.Println(nodes)
	//
	////root.print()
	//root.Right.Left.setValue(4)//node是指针后，在传递过程中，节点也会变为指针
	//root.Right.Left.print()//如果方法的node不带星，因为传的是值，所以改不掉
	//fmt.Println()
	//
	//pRoot := &root;
	//pRoot.print()
	//pRoot.setValue(222)
	//pRoot.print()
	//
	//var nRoot *Node
	//nRoot.setValue(200)
	//nRoot = pRoot
	//nRoot.setValue(300)
	//nRoot.print()

}

