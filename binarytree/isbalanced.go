package main

import (
	"fmt"
	"math"
)

/*
	是否是平衡二叉树,一颗二叉树的左右子


*/

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	fmt.Println(countFloor(root.Left) - countFloor(root.Right))
	if int(math.Abs(float64(countFloor(root.Left))-float64(countFloor(root.Right)))) > 1 {
		return false
	} else {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
}

func countFloor(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(countFloor(root.Left)), float64(countFloor(root.Right)))) + 1
}

func main() {

}
