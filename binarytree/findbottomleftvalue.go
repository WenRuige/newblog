package main

import (
	"github.com/newblog/binarytree/base"
)

func findBottomLeftValue(root *base.TreeNode) int {

	if root == nil {
		return 0
	}
	n := 1
	result := []int{}
	// 借助队列实现
	queue := []*base.TreeNode{root}
	for len(queue) > 0 {
		n--
		root := queue[0]
		queue = queue[1:]

		result = append(result, root.Val)
		if root.Left != nil {
			queue = append(queue, root.Left)
		}

		if root.Right != nil {
			queue = append(queue, root.Right)
		}

		if n == 0 {
			n = len(queue)
		}

	}

	return 0

}

func main() {

	treeNode := new(base.TreeNode)
	treeNode.Val = 1
	treeNode.Left = new(base.TreeNode)
	treeNode.Left.Val = 2
	treeNode.Right = new(base.TreeNode)
	treeNode.Right.Val = 3

	findBottomLeftValue(treeNode)
}
