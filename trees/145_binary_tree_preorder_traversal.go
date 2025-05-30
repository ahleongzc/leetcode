package trees

func PostorderTraversal(root *TreeNode) []int {
	values := make([]int, 0)

	var dfs func(root *TreeNode) 
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		dfs(root.Right)
		values = append(values, root.Val)
	}

	dfs(root)

	return values
}