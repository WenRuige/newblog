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

// 查看是否是子树
func issubtree(node *base.TreeNode, node2 *base.TreeNode) bool {

	if isTheSame(node, node2) {
		return true
	}

	return issubtree(node.Left, node2) && issubtree(node.Right, node2)
}

// 是否相同
func isTheSame(root *base.TreeNode, root2 *base.TreeNode) bool {
	if root == nil {
		return false
	}

	if root2 == nil {
		return true
	}

	if root.Val != root2.Val {
		return false
	}
	return isTheSame(root.Left, root2.Left) && isTheSame(root.Right, root2.Right)
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
	treeNode2.Val = 1
	treeNode2.Left = new(base.TreeNode)
	treeNode2.Left.Val = 2

	flag := issubtree(treeNode, treeNode2)
	fmt.Println(flag)
}
