package main

import (
	"fmt"

	"github.com/newblog/binarytree/base"
)

// 二叉树的最大深度
func maxDepth(root *base.TreeNode) int {

	if root == nil {
		return 0
	}
	leftLength := maxDepth(root.Left)
	rightLength := maxDepth(root.Right)
	if leftLength > rightLength {
		return leftLength + 1
	} else {
		return rightLength + 1
	}
}

// 基于非递归遍历二叉树最大深度
func maxDepthV2(root *base.TreeNode) int {

	if root == nil {
		return 0
	}
	queue := []*base.TreeNode{root}
	n := 1
	level := 0
	for len(queue) > 0 {
		n--
		root := queue[0]
		queue = queue[1:]
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
		if n == 0 {
			n = len(queue)
			level++

		}

	}
	return level
}
func main() {
	treeNode := new(base.TreeNode)
	treeNode.Val = 1
	treeNode.Left = new(base.TreeNode)
	treeNode.Left.Val = 2
	treeNode.Right = new(base.TreeNode)
	treeNode.Right.Val = 3
	res := maxDepthV2(treeNode)
	fmt.Println(res)
}
