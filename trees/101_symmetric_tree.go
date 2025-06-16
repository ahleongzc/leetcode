package trees

func IsSymmetric(root *TreeNode) bool {
	var dfs func(left, right *TreeNode)bool 
	dfs = func(left, right *TreeNode) bool {
		if left == nil && right == nil{
			return true
		}

		if left == nil || right == nil {
			return false
		}

		return left.Val == right.Val && dfs(left.Right, right.Left) && dfs(left.Left, right.Right)
	}

	if root == nil {
		return true
	}

	return dfs(root.Left, root.Right)
}
