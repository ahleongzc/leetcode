package trees

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var dfs func(root *TreeNode, currSum int) bool
	dfs = func(root *TreeNode, currSum int) bool {
		if root.Left == nil && root.Right == nil && currSum == targetSum {
			return true
		}

		left := false
		if root.Left != nil {
			left = dfs(root.Left, currSum+root.Left.Val)
		}

		right := false
		if root.Right != nil {
			right = dfs(root.Right, currSum+root.Right.Val)
		}

		return left || right
	}

	return dfs(root, root.Val)
}
