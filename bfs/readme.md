#### 广度优先算法(BFS)




##### 树的广度优先搜索算法
广度优先搜索算法`(Breadth-First-Search,BFS)`又称作宽度优先搜索,`BFS`是从根节点开始,沿着树的宽度来遍历树的节点,如果所有节点都被访问,则算法种种


* 首先将根节点放入队列中
* 从队列中取出第一个节点,验证是否是目标节点


###### 复杂度
* 由于所有节点都必须被存储,因此`BFS`空间复杂度为`O|V| + O|E|`,其中`V`是节点的数目,`E`是边的数目



eg: 让我们来看一个寻找每层最大的树的节点的值
```
func bfs(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	// 声明一个存储空间
	queue := make([]*TreeNode, 1, 1024)
	// 0为根节点
	queue[0] = root
	// n = 1 二叉树层数
	n := 1
	max := math.MinInt32
	for len(queue) > 0 {
		n--
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		if max < node.Val {
			max = node.Val
		}
		// 层数
		if n == 0 {
			n = len(queue)
			result = append(result, max)
			max = math.MinInt32
		}
	}
	return result
}

```