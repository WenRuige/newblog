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
======================================
		1
	2		3				2
4      5  				4		5
======================================

类似于第二种结构是不行的


Q:如何去思考二叉树递归过程


*/

type TreeNode = base.TreeNode

// 查看是否是子树
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if isSame(s, t) {
		return true
	}

	if s == nil { // 如果t的树大于s的树
		return false
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return t == nil
	}
	if t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}

func main() {
	treeNode := new(TreeNode)
	treeNode.Val = 1
	treeNode.Left = new(TreeNode)
	treeNode.Left.Val = 2
	treeNode.Right = new(TreeNode)
	treeNode.Right.Val = 3
	treeNode.Left.Left = new(TreeNode)
	treeNode.Left.Left.Val = 4
	treeNode.Left.Right = new(TreeNode)
	treeNode.Left.Right.Val = 5

	treeNode2 := new(TreeNode)
	treeNode2.Val = 2
	treeNode2.Left = new(TreeNode)
	treeNode2.Left.Val = 4
	treeNode2.Right = new(TreeNode)
	treeNode2.Right.Val = 5

	flag := isSubtree(treeNode, treeNode2)
	fmt.Println(flag)
}
