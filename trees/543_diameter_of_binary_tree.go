package trees

func DiameterOfBinaryTree(root *TreeNode) int {
	// The first element represents the current maximum diameter
	// The second element represents the height of the tree

	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		leftDiameter, leftHeight := dfs(node.Left)
		rightDiameter, rightHeight := dfs(node.Right)

		maxDiameter := max(
			leftDiameter,
			rightDiameter,
			leftHeight+rightHeight,
		)

		maxDepth := max(leftHeight, rightHeight) + 1

		return maxDiameter, maxDepth
	}

	maxDiameter, _ := dfs(root)

	return maxDiameter
}
