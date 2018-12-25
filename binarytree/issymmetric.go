package main

import (
	"fmt"

	"github.com/newblog/binarytree/base"
)

/*

 对称二叉树
	和相同二叉树是镜像问题.


*/
func isSymmetric(root *base.TreeNode) bool {
	if root == nil {
		return true
	}
	return checkSame(root.Left, root.Right)
}

// 判断两个数组是否相同
func checkSame(p *base.TreeNode, q *base.TreeNode) bool {

	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return checkSame(p.Left, q.Right) && checkSame(p.Right, q.Left)
}

func main() {
	treeNode := new(base.TreeNode)
	treeNode.Val = 1
	treeNode.Left = new(base.TreeNode)
	treeNode.Left.Val = 2
	treeNode.Right = new(base.TreeNode)
	treeNode.Right.Val = 2
	flag := isSymmetric(treeNode)
	fmt.Println(flag)

}
