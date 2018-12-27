package main

import "github.com/newblog/binarytree/base"

func findBottomLeftValue(root *base.TreeNode) int {

	if root == nil {
		return 0
	}
	// 借助队列实现
	queue := []*base.TreeNode{root}
	for len(queue) > 0 {
		root := queue[0]
		queue = queue[1:]
		if root.Left != nil {
			queue = append(queue, root.Left)
		}

		if root.Right != nil {
			queue = append(queue, root.Right)
		}

	}

}
