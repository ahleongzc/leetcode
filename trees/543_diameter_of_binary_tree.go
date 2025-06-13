package trees

func DiameterOfBinaryTree(root *TreeNode) int {
	// The first element represents the current maximum diameter
	// The second element represents the height of the tree

	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		leftDiameter, leftDepth := dfs(node.Left)
		rightDiameter, rightDepth := dfs(node.Right)

		maxDiameter := max(
			leftDiameter,
			rightDiameter,
			leftDepth+rightDepth,
		)

		maxDepth := max(leftDepth, rightDepth) + 1

		return maxDiameter, maxDepth
	}

	maxDiameter, _ := dfs(root)

	return maxDiameter
}
