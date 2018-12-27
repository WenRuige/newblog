package main

import (
	"fmt"

	"github.com/newblog/binarytree/base"
)

func flatten(root *base.TreeNode) {
	result := []int{}
	preTraversal(root, &result)
	fmt.Println(result)

}

// 前序遍历
func preTraversal(node *base.TreeNode, result *[]int) {
	*result = append(*result, node.Val)
	if node.Left != nil {
		preTraversal(node.Left, result)
	}
	if node.Right != nil {
		preTraversal(node.Right, result)
	}

}

func main() {
	treeNode := new(base.TreeNode)
	treeNode.Val = 1
	treeNode.Left = new(base.TreeNode)
	treeNode.Left.Val = 2
	treeNode.Right = new(base.TreeNode)
	treeNode.Right.Val = 3

	flatten(treeNode)
}
