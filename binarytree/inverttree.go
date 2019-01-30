package main

import "github.com/newblog/binarytree/base"

// 翻转二叉树
func invertTree(root *base.TreeNode) *base.TreeNode {

	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	root.Left = invertTree(root.Right)
	root.Right = invertTree(root.Left)
	return root

}

func main() {
	t := new(base.TreeNode)
	t.Val = 1
	t.Left = new(base.TreeNode)
	t.Left.Val = 2
}
