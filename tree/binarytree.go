// Package tree
// @Description:
package tree

//
//  Node
//  @Description: 二叉树的节点
//
type Node struct {
	Left  *Node
	Right *Node
	data  int
}

//
//  BinaryTree
//  @Description: 二叉树
//
type BinaryTree struct {
	root     *Node
	numNodes int
}

//
//  NewBinaryTree
//  @Description: 根据层序遍历创建二叉树
//  @receiver *BinaryTree
//  @param inputs
//
func (*BinaryTree) NewBinaryTree(inputs []string) {

}

func (this *BinaryTree) Inorder() []int {
	res := make([]int, 0)
	this.inorder(this.root, &res)
	return res
}

//
//  inorder
//  @Description: 中序遍历
//  @receiver this
//  @param root
//  @param res
//
func (this *BinaryTree) inorder(root *Node, res *[]int) {
	if root != nil {
		this.inorder(root.Left, res)
		*res = append(*res, root.data)
		this.inorder(root.Right, res)
	}
}
