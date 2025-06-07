package trees

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	var dfs func(p *TreeNode, q *TreeNode) bool
	dfs = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if p == nil || q == nil {
			return false
		}

		if p.Val != q.Val {
			return false
		}

		return dfs(p.Left, q.Left) && dfs(p.Right, q.Right)
	}

	return dfs(p, q)
}
