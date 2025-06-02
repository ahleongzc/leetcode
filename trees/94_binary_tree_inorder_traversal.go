package trees

func InorderTraversal(root *TreeNode) []int {
	values := make([]int, 0)

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		values = append(values, root.Val)
		dfs(root.Right)
	}

	dfs(root)

	return values
}
