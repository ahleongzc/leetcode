package trees

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	lowerNode := p
	higherNode := q

	if p.Val > q.Val {
		lowerNode = q
		higherNode = p
	}

	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}

		if root.Val >= lowerNode.Val && root.Val <= higherNode.Val {
			return root
		}

		if root.Val < lowerNode.Val {
			return dfs(root.Right)
		}

		return dfs(root.Left)
	}

	return dfs(root)
}
