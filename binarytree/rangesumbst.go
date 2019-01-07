package main

import (
	"fmt"

	"github.com/newblog/binarytree/base"
)

// 二叉树的范围和
/*
	二叉搜索树的概念
	如果根节点的左子树不为空,左子树始终小于根节点
	如果根节点的右子树不为空,右子树始终大于根节点

		4
	3		5


*/
func rangeSumBST(root *base.TreeNode, L int, R int) int {
	// 	根节点为空
	sum := 0
	if root == nil {
		return sum
	}
	if root.Val >= L && root.Val <= R {
		sum = root.Val
	}
	return rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R) + sum

}

func rangeSumBST2(root *base.TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	// 根节点小于右的话,说明需要向左遍历
	if root.Val > R {
		return rangeSumBST2(root.Left, L, R)

	} else if root.Val < L {
		return rangeSumBST2(root.Right, L, R)
	} else {
		return root.Val + rangeSumBST2(root.Left, L, R) + rangeSumBST2(root.Right, L, R)
	}

}

func main() {

	treeNode := new(base.TreeNode)
	treeNode.Val = 1
	treeNode.Left = new(base.TreeNode)
	treeNode.Left.Val = 2
	treeNode.Right = new(base.TreeNode)
	treeNode.Right.Val = 3

	result := rangeSumBST(treeNode, 1, 2)
	fmt.Println(result)
}
