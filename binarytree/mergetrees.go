package main

import (
	"fmt"

	"github.com/newblog/binarytree/base"
)

func mergeTrees(t1 *base.TreeNode, t2 *base.TreeNode) *base.TreeNode {

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	return &base.TreeNode{Val: t1.Val + t2.Val, Left: mergeTrees(t1.Left, t2.Left), Right: mergeTrees(t1.Right, t2.Right)}
}

func main() {

	t := new(base.TreeNode)
	t.Val = 1
	t.Left = new(base.TreeNode)
	t.Left.Val = 2

	t1 := new(base.TreeNode)
	t1.Val = 3
	t1.Left = new(base.TreeNode)
	t1.Left.Val = 5

	res := mergeTrees(t, t1)

	bianli(res)

}

func bianli(tree *base.TreeNode) {
	if tree.Left != nil {
		bianli(tree.Left)
	}
	fmt.Println(tree.Val)
	if tree.Right != nil {
		bianli(tree.Right)
	}

}
