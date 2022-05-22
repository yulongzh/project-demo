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
