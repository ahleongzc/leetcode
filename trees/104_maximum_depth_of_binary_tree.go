package trees

func MaxDepth(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		return max(1+dfs(root.Left), 1+dfs(root.Right))
	}

	return dfs(root)
}
