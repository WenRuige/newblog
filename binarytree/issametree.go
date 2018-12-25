package main

import (
	"fmt"

	"github.com/newblog/binarytree/base"
)

//import "github.com/newblog/binarytree/base"

//type TreeNode = base.TreeNode

/*

	相同的树
		确保两棵树,左右子树都相等

*/
func isSameTree(p *base.TreeNode, q *base.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	treeNode := new(base.TreeNode)
	treeNode.Val = 1
	treeNode.Left = new(base.TreeNode)
	treeNode.Left.Val = 2
	treeNode.Right = new(base.TreeNode)
	treeNode.Right.Val = 3

	treeNode2 := new(base.TreeNode)
	treeNode2.Val = 1
	treeNode2.Left = new(base.TreeNode)
	treeNode2.Left.Val = 2
	treeNode2.Right = new(base.TreeNode)
	treeNode2.Right.Val = 3

	flag := isSameTree(treeNode, treeNode2)
	fmt.Println(flag)
}
