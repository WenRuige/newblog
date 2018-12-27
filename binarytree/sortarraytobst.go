package main

import (
	"fmt"

	"github.com/newblog/binarytree/base"
)

/*
	二叉搜索树
 		- 或者它是一颗空树
		- 若左子树不空,则左子树上所有的节点都小于它的根节点的值
		- 若右子树不空,则右子树上所有的节点都大于它的根节点的值



====== 貌似也是利用了一个trick来实现的 ========

*/
func sortedArrayToBST(nums []int) *base.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, low int, high int) *base.TreeNode {
	if low > high {
		return nil
	}
	mid := (low + high) / 2
	treeNode := new(base.TreeNode)

	treeNode.Val = nums[mid]
	treeNode.Left = helper(nums, low, mid-1)
	treeNode.Right = helper(nums, mid+1, high)
	return treeNode
}

func main() {

	res := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	fmt.Println(res)

}
