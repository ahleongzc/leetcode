package trees

func PreorderTraversal(root *TreeNode) []int {
	values := make([]int, 0)

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		values = append(values, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)

	return values
}
