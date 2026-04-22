package trees

func rightSideView(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue = append(queue, root)

	for len(queue) != 0 {
		last := queue[len(queue)-1]
		res = append(res, last.Val)

		currQueueLen := len(queue)
		for range currQueueLen {
			curr := queue[0]
			queue = queue[1:]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}

	return res
}

func rightSideViewDFS(root *TreeNode) []int {
	res := make([]int, 0)

	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(res) == depth {
			res = append(res, root.Val)
		}

		dfs(root.Right, depth+1)
		dfs(root.Left, depth+1)
	}

	dfs(root, 0)
	return res
}
