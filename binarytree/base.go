package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 遍历二叉树所有的值

func traversingBinaryTree(treeNode TreeNode) {

	//if treeNode != nil {
	//
	//}

}

func recursiveTest(k int) int {
	if k == 0 {
		return 0
	} else {
		return recursiveTest(k-1) + k
	}
}
func main() {
	//t := new(TreeNode)
	//t.Val = 1
	//t.Left = new(TreeNode)
	//t.Left.Val = 2
	//t.Right = new(TreeNode)
	//t.Right.Val = 3

	fmt.Println(recursiveTest(5))

}
