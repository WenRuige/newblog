package base

//import (
//	"fmt"
//	"math"
//)
//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//
//// 遍历二叉树所有的值
//
//func traversingBinaryTree(treeNode TreeNode) {
//
//	//if treeNode != nil {
//	//
//	//}
//
//}
//
//func recursiveTest(k int) int {
//	if k == 0 {
//		return 0
//	} else {
//		return recursiveTest(k-1) + k
//	}
//}
//
//// 层次遍历二叉树
//func levelTree(root *TreeNode) []int {
//	var result []int
//	queue := []*TreeNode{root}
//
//	for len(queue) > 0 {
//		root = queue[0]
//		queue = queue[1:]
//		result = append(result, root.Val)
//		// 先把 Right 边的节点放入 queue 才能保证
//		// 最后剩下的 root 是 bottom left
//		if root.Right != nil {
//			queue = append(queue, root.Right)
//		}
//		if root.Left != nil {
//			queue = append(queue, root.Left)
//		}
//	}
//
//	return result
//}
//
//func getTreeDepthRecursion(node *TreeNode) int {
//	if node == nil {
//		return 0
//	}
//	left := getTreeDepthRecursion(node.Left)
//	right := getTreeDepthRecursion(node.Right)
//	return int(math.Max(float64(left), float64(right))) + 1
//}
//
//func getTreeDepthNonRecursion(root *TreeNode) int {
//	// 层序遍历求深度
//	queue := make([]*TreeNode, 1, 1024)
//	queue[0] = root
//	n := 1
//	level := 1
//	if len(queue) > 0 {
//		n--
//		node := queue[0]
//		queue = queue[1:]
//		if node.Left != nil {
//			queue = append(queue, node.Left)
//		}
//		if node.Right != nil {
//			queue = append(queue, node.Right)
//		}
//		if n == 0 {
//			level++
//		}
//
//	}
//	return level
//}
//
//// 构建满二叉树 数组 -> tree
//
//func NewTree() {
//	var node *TreeNode
//	node = generateBinaryTree(node, 123123)
//	node = generateBinaryTree(node, 1)
//	node = generateBinaryTree(node, 2)
//	node = generateBinaryTree(node, 3)
//	node = generateBinaryTree(node, 5)
//
//	res := levelTree(node)
//
//	fmt.Println(res)
//}
//
//// 生成binary Tree
//func generateBinaryTree(node *TreeNode, num int) *TreeNode {
//	if node == nil {
//		return &TreeNode{num, nil, nil}
//	}
//	if node.Val < num {
//		node.Left = generateBinaryTree(node.Left, num)
//	} else {
//		node.Right = generateBinaryTree(node.Right, num)
//	}
//	return node
//}
//
//// 构建平衡二叉树
//
//func main() {
//
//	NewTree()
//	//t := new(TreeNode)
//	//t.Val = 1
//	//t.Left = new(TreeNode)
//	//t.Left.Val = 2
//	//t.Right = new(TreeNode)
//	//t.Right.Val = 3
//	//
//	//res := getTreeDepthNonRecursion(t)
//	//fmt.Println(res)
//
//	//fmt.Println(recursiveTest(5))
//
//}
