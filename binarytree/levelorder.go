package main

import (
	"fmt"

	"github.com/newblog/binarytree/base"
)

/*

	层序遍历二叉树要点:
		需要借助一个队列(由于队列是先进先出的),所以可以先让根如队,在根出队的时候同时打印根
		并让根的左孩子入队,然后在让根的右孩子入队,依次类推.



*/
// 层序遍历二叉树
func levelOrder(root *base.TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}
	queue := []*base.TreeNode{root}
	n := 1
	res := []int{}
	final := [][]int{}
	for len(queue) > 0 {
		n--
		root = queue[0]
		queue = queue[1:]
		res = append(res, root.Val)
		if root.Left != nil {
			queue = append(queue, root.Left)
		}

		if root.Right != nil {
			queue = append(queue, root.Right)
		}
		if n == 0 {
			n = len(queue)
			final = append(final, res)
			res = []int{}
		}
	}

	return final
}

func levelOrderV2(root *base.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*base.TreeNode{root}
	res := []int{}
	final := [][]int{}
	n := 1
	for len(queue) != 0 {
		n--
		root := queue[0]
		queue = queue[1:]

		res = append(res, root.Val)
		if root.Left != nil {
			queue = append(queue, root.Left)
		}

		if root.Right != nil {
			queue = append(queue, root.Right)
		}

		if n == 0 {
			n = len(queue)
			final = append(final, res)
			res = []int{}
		}
	}
	newFinal := [][]int{}
	for i := len(final) - 1; i >= 0; i-- {
		newFinal = append(newFinal, final[i])
	}
	return newFinal

}

func main() {
	treeNode := new(base.TreeNode)
	treeNode.Val = 1
	treeNode.Left = new(base.TreeNode)
	treeNode.Left.Val = 2
	treeNode.Right = new(base.TreeNode)
	treeNode.Right.Val = 3

	res := levelOrderV2(treeNode)
	fmt.Println(res)
}
