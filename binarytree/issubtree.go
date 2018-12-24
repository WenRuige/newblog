package main

import (
	"fmt"

	"github.com/newblog/binarytree/base"
)

/*

判断一颗二叉树是不是另外一颗二叉树的子树

临界点:
     - 第一个二叉树为空,第二个二叉树还不为空 return false
	 - 第二个二叉树为空,return true


		1						1
	2		3				2		3
4				5




*/

func issubtree(node *base.TreeNode, node2 *base.TreeNode) bool {
	if node == nil {
		return false
	}

	if node2 == nil {
		return true
	}

	if node.Val != node2.Val {
		return false
	}

	return issubtree(node.Left, node2.Left) && issubtree(node.Right, node2.Right)

}
func main() {
	treeNode := new(base.TreeNode)
	treeNode.Val = 1
	treeNode.Left = new(base.TreeNode)
	treeNode.Left.Val = 2
	treeNode.Right = new(base.TreeNode)
	treeNode.Right.Val = 3
	treeNode.Left.Left = new(base.TreeNode)
	treeNode.Left.Left.Val = 4
	treeNode.Left.Right = new(base.TreeNode)
	treeNode.Left.Right.Val = 5

	treeNode2 := new(base.TreeNode)
	treeNode2.Val = 3

	flag := issubtree(treeNode, treeNode2)
	fmt.Println(flag)
}
