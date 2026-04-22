package trees

func goodNodes(root *TreeNode) int {
	res := 0
	var dfs func(root *TreeNode, prev int)
	dfs = func(root *TreeNode, prev int) {
		if root == nil {
			return
		}
		if root.Val >= prev {
			res++
		}

		next := max(root.Val, prev)
		dfs(root.Left, next)
		dfs(root.Right, next)
	}

	dfs(root, root.Val-1)
	return res
}
